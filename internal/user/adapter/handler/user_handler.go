package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"hostel-service/internal/user/domain"
	"hostel-service/internal/user/port"
	"hostel-service/internal/user/service"
	"hostel-service/pkg/util"
)

func NewUserHandler(
	service service.UserService,
	validate *validator.Validate,
) port.UserHandler {
	return &HttpUserHandler{
		service:  service,
		validate: validate,
	}
}

type HttpUserHandler struct {
	service  service.UserService
	validate *validator.Validate
}

func (h *HttpUserHandler) Login(w http.ResponseWriter, r *http.Request) {
	input := &domain.LoginReq{}
	err := json.NewDecoder(r.Body).Decode(input)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}

	user, token, code, err := h.service.Login(r.Context(), input.Username, input.Password)
	if code == -1 {
		util.JsonInternalError(w, err)
		return
	} else if code == 0 {
		util.JsonBadRequest(w, err, "Tài khoản hoặc mật khẩu không chính xác")
		return
	}

	res := domain.LoginRes{
		Token:   token,
		Profile: user,
	}
	if code == 2 {
		res.IsResetPass = true
	}

	util.JsonOK(w, res)
}

func (h *HttpUserHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := &domain.RegisterReq{}
	err := json.NewDecoder(r.Body).Decode(user)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}

	err = h.service.Register(r.Context(), user.Username, user.Phone, user.Name)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpUserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	res, err := h.service.GetUserProfile(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w, res)
}

func (h *HttpUserHandler) SearchRoommates(w http.ResponseWriter, r *http.Request) {
	filter := &domain.RoommateFilter{
		Sort: "created_at desc",
	}
	err := json.NewDecoder(r.Body).Decode(filter)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}

	res, total, err := h.service.SearchRoommates(r.Context(), filter)
	if err != nil {
		util.JsonInternalError(w, err)
	}
	util.JsonOK(w, res, total)
}

func (h *HttpUserHandler) GetRoommateById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	user, err := h.service.GetRoommateById(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else if user == nil {
		util.Json(w, http.StatusNotFound, util.Response{})
	} else {
		util.JsonOK(w, user)
	}
}

func (h *HttpUserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	req := &domain.UpdatePasswordReq{}
	userId := r.Context().Value("userId").(string)
	err := json.NewDecoder(r.Body).Decode(req)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}

	err = h.service.UpdatePassword(r.Context(), userId, req.OldPassword, req.NewPassword)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpUserHandler) VerifyPhone(w http.ResponseWriter, r *http.Request) {
	req := &domain.VerifyPhoneReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}
	userId := r.Context().Value("userId").(string)

	err = h.service.VerifyPhone(r.Context(), userId, req.Phone)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpUserHandler) VerifyPhoneOTP(w http.ResponseWriter, r *http.Request) {
	req := &domain.VerifyOTPReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}
	userId := r.Context().Value("userId").(string)

	err = h.service.VerifyPhoneOTP(r.Context(), userId, req.OTP)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

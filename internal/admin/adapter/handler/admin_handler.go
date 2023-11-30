package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"hostel-service/internal/admin/domain"
	"hostel-service/internal/admin/port"
	"hostel-service/internal/admin/service"
	"hostel-service/pkg/util"
)

func NewAdminHandler(
	service service.AdminService,
	validate *validator.Validate,
) port.AdminHandler {
	return &HttpAdminHandler{
		service:  service,
		validate: validate,
	}
}

type HttpAdminHandler struct {
	service  service.AdminService
	validate *validator.Validate
}

func (h *HttpAdminHandler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := &domain.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(credentials)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}

	admin, er2 := h.service.GetByUsername(r.Context(), credentials.Username)
	if er2 != nil {
		util.JsonInternalError(w, errors.New("internal server error"))
		return
	}
	if admin == nil || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(credentials.Password)) != nil {
		util.Json(w, http.StatusNotFound, util.Response{
			Status:  "admin not match",
			Message: "Tài khoản hoặc mật khẩu không chính xác",
		})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  admin.Id,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})
	token, er3 := claims.SignedString(domain.ADMIN_SECRET_KEY)
	if er3 != nil {
		util.JsonInternalError(w, errors.New("internal server error"))
		return
	}
	res := domain.LoginResponse{
		Token:   token,
		Profile: nil,
	}
	util.JsonOK(w, res)

}

func (h *HttpAdminHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	req := &domain.UpdatePasswordRequest{}
	adminId := r.Context().Value("adminId").(string)
	err := json.NewDecoder(r.Body).Decode(req)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}
	err = h.service.UpdatePassword(r.Context(), adminId, req.OldPassword, req.NewPassword)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpAdminHandler) GetAdminProfile(w http.ResponseWriter, r *http.Request) {
	adminId := r.Context().Value("adminId").(string)
	res, err := h.service.GetAdminProfile(r.Context(), adminId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w, res)
}

func (h *HttpAdminHandler) SearchPosts(w http.ResponseWriter, r *http.Request) {
	post := &domain.PostFilter{
		Sort: "created_at desc",
	}
	err := json.NewDecoder(r.Body).Decode(post)
	defer r.Body.Close()
	if err != nil {
		util.Json(w, http.StatusBadRequest, util.Response{Status: err.Error()})
		return
	}
	posts, total, err := h.service.SearchPosts(r.Context(), post)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.Json(w, http.StatusOK, util.Response{
		Data:  posts,
		Total: total,
	})
}

func (h *HttpAdminHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	post, err := h.service.GetPostById(r.Context(), postId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	if post == nil {
		util.Json(w, http.StatusNotFound, util.Response{Status: "not found"})
		return
	}
	util.JsonOK(w, post)
}

func (h *HttpAdminHandler) UpdatePostStatus(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	var status string
	switch path.Base(r.URL.Path) {
	case "disable":
		status = "I"
	case "active":
		status = "A"
	case "verify":
		status = "V"
	}
	_, err := h.service.UpdatePostStatus(r.Context(), postId, status)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

// User
func (h *HttpAdminHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	filter := &domain.UserFilter{
		Sort: "created_at desc",
	}
	err := json.NewDecoder(r.Body).Decode(filter)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}

	res, total, err := h.service.SearchUsers(r.Context(), filter)
	if err != nil {
		util.JsonInternalError(w, err)
	}
	util.JsonOK(w, res, total)
}

func (h *HttpAdminHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	user, err := h.service.GetUserById(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else if user == nil {
		util.Json(w, http.StatusNotFound, util.Response{})
	} else {
		util.JsonOK(w, user)
	}
}

func (h *HttpAdminHandler) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	status := ""
	switch path.Base(r.URL.Path) {
	case "disable":
		status = "I"
	case "active":
		status = "A"
	}

	err := h.service.UpdateUserStatus(r.Context(), userId, status)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpAdminHandler) ResetUserPassword(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	err := h.service.ResetPassword(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

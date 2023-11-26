package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"hostel-service/internal/admin/domain"
	"hostel-service/internal/admin/service"
	"hostel-service/internal/package/util"
)

func NewAdminHandler(
	service service.AdminService,
	validate *validator.Validate,
) *HttpAdminHandler {
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
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: err.Error(),
		})
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
			Message: "Adminname or password not match",
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
	util.Json(w, http.StatusOK, res)

}

func (h *HttpAdminHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	req := &domain.UpdatePasswordRequest{}
	adminId := r.Context().Value("userId").(string)
	err := json.NewDecoder(r.Body).Decode(req)
	defer r.Body.Close()
	if err != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: err.Error(),
		})
		return
	}
	err = h.service.UpdatePassword(r.Context(), adminId, req.OldPassword, req.NewPassword)
	if err != nil {
		util.Json(w, http.StatusInternalServerError, util.Response{Status: err.Error()})
		return
	}
	util.JsonOK(w)
}

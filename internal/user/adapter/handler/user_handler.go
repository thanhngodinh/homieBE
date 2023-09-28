package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/core-go/core"
	sv "github.com/core-go/core"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"hostel-service/internal/user/domain"
	"hostel-service/internal/user/service"
	"hostel-service/internal/util"
)

func NewUserHandler(service service.UserService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *HttpUserHandler {
	return &HttpUserHandler{service: service, validate: validate, logError: logError}
}

type HttpUserHandler struct {
	service  service.UserService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpUserHandler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := &domain.LoginRequest{}
	er1 := core.Decode(w, r, credentials)
	if er1 == nil {
		user, er2 := h.service.GetByUsername(r.Context(), credentials.Username)
		if er2 != nil {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
			return
		}
		if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
			core.Respond(w, r, http.StatusNotFound, nil, nil, nil, nil)
			return
		}
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Audience:  user.Id,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
		token, er3 := claims.SignedString(domain.SECRET_KEY)
		if er3 != nil {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
			return
		}
		res := domain.LoginResponse{
			Token:   token,
			Profile: nil,
		}
		JSON(w, http.StatusOK, res)
	}
}

func (h *HttpUserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	er1 := core.Decode(w, r, &user)
	if er1 == nil {
		errors, er2 := h.validate(r.Context(), &user)
		if errors != nil || er2 != nil {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		}
		user.Id = uuid.New().String()
		hashedPassword, er3 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if er3 != nil {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)
		now := time.Now()
		user.CreatedAt = &now
		res, er3 := h.service.Create(r.Context(), &user)
		JSON(w, http.StatusOK, util.Response{
			Data: res,
		})
	}
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

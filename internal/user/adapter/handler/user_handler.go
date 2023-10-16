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
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"hostel-service/internal/package/util"
	"hostel-service/internal/user/domain"
	"hostel-service/internal/user/service"
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
		util.Json(w, http.StatusOK, res)
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
		err := h.service.Create(r.Context(), &user)
		if err != nil {
			util.Json(w, http.StatusInternalServerError, util.Response{Status: err.Error()})
			return
		}
		util.JsonOK(w)
	}
}

func (h *HttpUserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	req := &domain.UpdatePasswordRequest{}
	userId := r.Context().Value("userId").(string)
	err := core.Decode(w, r, req)
	if err != nil {
		util.Json(w, http.StatusBadRequest, util.Response{Status: err.Error()})
		return
	}
	err = h.service.UpdatePassword(r.Context(), userId, req.OldPassword, req.NewPassword)
	if err != nil {
		util.Json(w, http.StatusInternalServerError, util.Response{Status: err.Error()})
		return
	}
	util.JsonOK(w)
}

func (h *HttpUserHandler) SearchRoommates(w http.ResponseWriter, r *http.Request) {
	filter := &domain.RoommateFilter{
		Sort: "created_at desc",
	}
	err := json.NewDecoder(r.Body).Decode(filter)
	defer r.Body.Close()
	if err != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: err.Error(),
		})
		return
	}
	res, total, err := h.service.SearchRoommates(r.Context(), filter)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  res,
			Total: total,
		})
	}
}

func (h *HttpUserHandler) GetRoommateById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
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

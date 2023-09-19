package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

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

func (h *HttpUserHandler) GetPostLikedByUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	res, err := h.service.GetPostLikedByUser(r.Context(), userId)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else {
		JSON(w, http.StatusOK, util.Response{
			Data:  res.Data,
			Total: res.Total,
		})
	}
}

func (h *HttpUserHandler) UserLikePost(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if len(userId) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorUserIdEmpty.Error(),
		})
		return
	}
	postId := mux.Vars(r)["postId"]
	if len(userId) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorPostIdEmpty.Error(),
		})
		return
	}

	res, err := h.service.UserLikePost(r.Context(), userId, postId)
	if err != nil {
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else {
		if res == 1 {
			JSON(w, http.StatusOK, util.Response{
				Data: fmt.Sprintf("update like state successfully"),
			})
		} else {
			JSON(w, http.StatusNotFound, util.Response{
				Data: fmt.Sprintf("not found"),
			})
		}
	}
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

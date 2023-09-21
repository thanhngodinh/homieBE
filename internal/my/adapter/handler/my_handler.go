package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

	"hostel-service/internal/my/service"
	"hostel-service/internal/util"
)

func NewMyHandler(service service.MyService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *HttpMyHandler {
	return &HttpMyHandler{service: service, validate: validate, logError: logError}
}

type HttpMyHandler struct {
	service  service.MyService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpMyHandler) GetMyPostLiked(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	res, err := h.service.GetMyPostLiked(r.Context(), userId)
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

func (h *HttpMyHandler) GetMyPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	res, err := h.service.GetMyPosts(r.Context(), userId)
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

func (h *HttpMyHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	postId := mux.Vars(r)["postId"]
	if len(userId) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorPostIdEmpty.Error(),
		})
		return
	}

	res, err := h.service.LikePost(r.Context(), userId, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

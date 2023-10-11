package handler

import (
	"context"
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
	userId := r.Context().Value("userId").(string)
	res, err := h.service.GetMyPostLiked(r.Context(), userId)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  res.Data,
			Total: res.Total,
		})
	}
}

func (h *HttpMyHandler) GetMyPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	res, err := h.service.GetMyPosts(r.Context(), userId)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  res.Data,
			Total: res.Total,
		})
	}
}

func (h *HttpMyHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorPostIdEmpty.Error(),
		})
		return
	}

	res, err := h.service.LikePost(r.Context(), userId, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if res == 1 {
			util.Json(w, http.StatusOK, util.Response{
				Data: fmt.Sprintf("update like state successfully"),
			})
		} else {
			util.Json(w, http.StatusNotFound, util.Response{
				Data: fmt.Sprintf("not found"),
			})
		}
	}
}

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

	"hostel-service/internal/my/domain"
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
		util.Json(w, http.StatusInternalServerError, util.Response{
			Status: err.Error(),
		})
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  res.Data,
			Total: res.Total,
		})
	}
}

func (h *HttpMyHandler) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	res, err := h.service.GetMyProfile(r.Context(), userId)
	if err != nil {
		h.logError(r.Context(), err.Error())
		util.Json(w, http.StatusInternalServerError, util.Response{
			Status: err.Error(),
		})
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data: res,
		})
	}
}

func (h *HttpMyHandler) UpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	var user domain.UpdateMyProfileReq
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: err.Error(),
		})
		return
	}
	errors, err := h.validate(r.Context(), &user)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		h.logError(r.Context(), err.Error())
		util.Json(w, http.StatusUnprocessableEntity, errors)
		return
	}
	err = h.service.UpdateMyProfile(r.Context(), &user)
	if err != nil {
		h.logError(r.Context(), err.Error())
		if util.IsDefinedErrorType(err) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: err.Error(),
			})
		} else {
			util.Json(w, http.StatusInternalServerError, util.Response{
				Status: err.Error(),
			})
		}
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data: user,
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
		util.Json(w, http.StatusInternalServerError, util.Response{
			Status: err.Error(),
		})
	} else {
		if res == 1 {
			util.JsonOK(w)
		} else {
			util.Json(w, http.StatusNotFound, util.Response{
				Status: fmt.Sprintf("not found"),
			})
		}
	}
}

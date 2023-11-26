package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"hostel-service/internal/my/domain"
	"hostel-service/internal/my/service"
	"hostel-service/package/util"
)

func NewMyHandler(service service.MyService, validate *validator.Validate) *HttpMyHandler {
	return &HttpMyHandler{service: service, validate: validate}
}

type HttpMyHandler struct {
	service  service.MyService
	validate *validator.Validate
}

func (h *HttpMyHandler) GetMyPostLiked(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	res, total, err := h.service.GetMyPostLiked(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
	}
	util.JsonOK(w, res, total)
}

func (h *HttpMyHandler) GetMyPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	res, total, err := h.service.GetMyPosts(r.Context(), userId)
	if err != nil {
		util.Json(w, http.StatusInternalServerError, util.Response{
			Status: err.Error(),
		})
	}
	util.JsonOK(w, res, total)
}

func (h *HttpMyHandler) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	res, err := h.service.GetMyProfile(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w, res)
}

func (h *HttpMyHandler) UpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	var user domain.UpdateMyProfileReq
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}
	user.Id = r.Context().Value("userId").(string)
	err = h.service.UpdateMyProfile(r.Context(), &user)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpMyHandler) UpdateMyAvatar(w http.ResponseWriter, r *http.Request) {
	var req domain.UpdateMyAvatarReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}
	req.Id = r.Context().Value("userId").(string)
	err = h.service.UpdateMyAvatar(r.Context(), &req)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)

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

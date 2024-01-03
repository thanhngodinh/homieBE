package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"hostel-service/internal/post/domain"
	"hostel-service/internal/post/service"
	"hostel-service/pkg/util"
)

func NewPostHandler(
	service service.PostService,
	validate *validator.Validate,
) *HttpPostHandler {
	return &HttpPostHandler{
		service:  service,
		validate: validate,
	}
}

type HttpPostHandler struct {
	service  service.PostService
	validate *validator.Validate
}

func (h *HttpPostHandler) SearchPosts(w http.ResponseWriter, r *http.Request) {
	post := &domain.PostFilter{
		Sort: "created_at desc",
	}
	er1 := json.NewDecoder(r.Body).Decode(post)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	userId := r.Context().Value("userId").(string)
	posts, total, err := h.service.SearchPosts(r.Context(), post, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  posts,
			Total: total,
		})
	}
}

func (h *HttpPostHandler) ElasticSearchPosts(w http.ResponseWriter, r *http.Request) {
	post := &domain.PostFilter{
		Sort: "created_at desc",
	}
	er1 := json.NewDecoder(r.Body).Decode(post)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	userId := r.Context().Value("userId").(string)
	posts, total, err := h.service.ESearchPosts(r.Context(), post, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  posts,
			Total: total,
		})
	}
}

func (h *HttpPostHandler) GetSuggestPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	posts, total, err := h.service.GetSuggestPosts(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  posts,
			Total: total,
		})
	}
}

func (h *HttpPostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	userId := r.Context().Value("userId").(string)
	post, err := h.service.GetPostById(r.Context(), code, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else if post == nil {
		util.Json(w, http.StatusNotFound, util.Response{Status: "not found"})
	} else {
		util.JsonOK(w, post)
	}
}

func (h *HttpPostHandler) GetCompare(w http.ResponseWriter, r *http.Request) {
	post1 := mux.Vars(r)["post1"]
	post2 := mux.Vars(r)["post2"]
	if len(post1) == 0 || len(post2) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	userId := r.Context().Value("userId").(string)
	post, err := h.service.GetCompare(r.Context(), post1, post2, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else if post == nil {
		util.Json(w, http.StatusNotFound, util.Response{})
	} else {
		util.JsonOK(w, domain.Compare{Post1: post[0], Post2: post[1]})
	}
}

func (h *HttpPostHandler) CheckCreatePost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	if userId == "" {
		util.Json(w, http.StatusOK, util.Response{
			Status:  "not login",
			Message: "User must login",
		})
		return
	}
	res, err := h.service.CheckCreatePost(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	if res == 0 {
		util.Json(w, http.StatusOK, util.Response{
			Status:  "not verify phone",
			Message: "User must verify phone",
		})
		return
	}
	util.JsonOK(w)
}

func (h *HttpPostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &domain.Post{}
	er1 := json.NewDecoder(r.Body).Decode(post)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	post.CreatedBy = r.Context().Value("userId").(string)
	_, er3 := h.service.CreatePost(r.Context(), post)
	if er3 != nil {
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			util.JsonInternalError(w, errors.New("internal server error"))
		}
	} else {
		util.Json(w, http.StatusCreated, util.Response{
			Data: post,
		})
	}
}

func (h *HttpPostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.UpdatePostReq
	er1 := json.NewDecoder(r.Body).Decode(&post)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	if len(post.Id) == 0 {
		post.Id = code
	} else if code != post.Id {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeNotMatch.Error(),
		})
		return
	}
	_, er3 := h.service.UpdatePost(r.Context(), &post)
	if er3 != nil {
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			util.JsonInternalError(w, errors.New("internal server error"))
		}
	} else {
		util.JsonOK(w, post)
	}
}

func (h *HttpPostHandler) HiddenPost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	_, err := h.service.HiddenPost(r.Context(), postId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpPostHandler) ActivePost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	_, err := h.service.ActivePost(r.Context(), postId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpPostHandler) ExtendPost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}

	_, err := h.service.ExtendPost(r.Context(), postId)
	if err != nil {
		util.JsonInternalError(w, err)
		return
	}
	util.JsonOK(w)
}

func (h *HttpPostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	res, err := h.service.DeletePost(r.Context(), code)
	if err != nil {
		util.JsonInternalError(w, errors.New("internal server error"))
		return
	}
	if res == 0 {
		util.Json(w, http.StatusNotFound, util.Response{
			Status: "not found",
		})
		return
	}
	util.JsonOK(w)
}

package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"hostel-service/internal/package/util"
	"hostel-service/internal/post/domain"
	"hostel-service/internal/post/service"
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

func (h *HttpPostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	pageIdxParam := r.URL.Query().Get("pageIdx")
	pageSizeParam := r.URL.Query().Get("pageSize")
	sort := r.URL.Query().Get("sort")

	hostel := &domain.PostFilter{
		Sort: "created_at desc",
	}
	if len(sort) > 0 {
		hostel.Sort = sort
	}
	if len(pageIdxParam) > 0 {
		pageIdx, err := strconv.Atoi(pageIdxParam)
		if err != nil {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: util.ErrorWrongTypePageIdx.Error(),
			})
			return
		}
		hostel.PageIdx = pageIdx
	} else {
		hostel.PageIdx = 0
	}

	if len(pageSizeParam) > 0 {
		pageSize, err := strconv.Atoi(pageSizeParam)
		if err != nil {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: util.ErrorWrongTypePageSize.Error(),
			})
			return
		}
		hostel.PageSize = pageSize
	} else {
		hostel.PageSize = 10
	}

	posts, total, err := h.service.GetPosts(r.Context(), hostel, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  posts,
			Total: total,
		})
	}
}

func (h *HttpPostHandler) SearchPosts(w http.ResponseWriter, r *http.Request) {
	hostel := &domain.PostFilter{
		Sort: "created_at desc",
	}
	er1 := json.NewDecoder(r.Body).Decode(hostel)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	userId := r.Context().Value("userId").(string)
	posts, total, err := h.service.GetPosts(r.Context(), hostel, userId)
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
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	userId := r.Context().Value("userId").(string)
	hostel, err := h.service.GetPostById(r.Context(), code, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else if hostel == nil {
		util.Json(w, http.StatusNotFound, util.Response{})
	} else {
		util.Json(w, http.StatusOK, hostel)
	}
}

func (h *HttpPostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Post
	er1 := json.NewDecoder(r.Body).Decode(&hostel)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	hostel.CreatedBy = r.Context().Value("userId").(string)
	_, er3 := h.service.CreatePost(r.Context(), &hostel)
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
			Data: hostel,
		})
	}
}

func (h *HttpPostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Post
	er1 := json.NewDecoder(r.Body).Decode(&hostel)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	if len(hostel.Id) == 0 {
		hostel.Id = code
	} else if code != hostel.Id {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeNotMatch.Error(),
		})
		return
	}
	_, er3 := h.service.UpdatePost(r.Context(), &hostel)
	if er3 != nil {
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			util.JsonInternalError(w, errors.New("internal server error"))
		}
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data: hostel,
		})
	}
}

func (h *HttpPostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	res, err := h.service.DeletePost(r.Context(), code)
	if err != nil {
		util.JsonInternalError(w, errors.New("internal server error"))
	} else {
		if res == 1 {
			util.Json(w, http.StatusOK, util.Response{
				Data: fmt.Sprintf("delete %s successfully", code),
			})
		} else {
			util.Json(w, http.StatusNotFound, util.Response{
				Data: fmt.Sprintf("not found %s", code),
			})
		}
	}
}

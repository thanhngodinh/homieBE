package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

	"hostel-service/internal/hostel/domain"
	"hostel-service/internal/hostel/service"
	"hostel-service/internal/package/util"
)

func NewHostelHandler(
	service service.HostelService,
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
	logError func(context.Context, string, ...map[string]interface{})) *HttpHostelHandler {
	return &HttpHostelHandler{
		service:  service,
		validate: validate,
		logError: logError}
}

type HttpHostelHandler struct {
	service  service.HostelService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpHostelHandler) GetHostels(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	pageIdxParam := r.URL.Query().Get("pageIdx")
	pageSizeParam := r.URL.Query().Get("pageSize")
	sort := r.URL.Query().Get("sort")

	hostel := &domain.HostelFilter{
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

	hostels, total, err := h.service.GetHostels(r.Context(), hostel, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  hostels,
			Total: total,
		})
	}
}

func (h *HttpHostelHandler) SearchHostels(w http.ResponseWriter, r *http.Request) {
	hostel := &domain.HostelFilter{
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
	hostels, total, err := h.service.GetHostels(r.Context(), hostel, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  hostels,
			Total: total,
		})
	}
}

func (h *HttpHostelHandler) GetSuggestHostels(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	hostels, total, err := h.service.GetSuggestHostels(r.Context(), userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data:  hostels,
			Total: total,
		})
	}
}

func (h *HttpHostelHandler) GetHostelById(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	userId := r.Context().Value("userId").(string)
	hostel, err := h.service.GetHostelById(r.Context(), code, userId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else if hostel == nil {
		util.Json(w, http.StatusNotFound, util.Response{})
	} else {
		util.Json(w, http.StatusOK, hostel)
	}
}

func (h *HttpHostelHandler) CreateHostel(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Hostel
	er1 := json.NewDecoder(r.Body).Decode(&hostel)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	errors, er2 := h.validate(r.Context(), &hostel)
	if er2 != nil {
		h.logError(r.Context(), er2.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		h.logError(r.Context(), er2.Error())
		util.Json(w, http.StatusUnprocessableEntity, errors)
		return
	}
	hostel.CreatedBy = r.Context().Value("userId").(string)
	_, er3 := h.service.CreateHostel(r.Context(), &hostel)
	if er3 != nil {
		h.logError(r.Context(), er3.Error())
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		}
	} else {
		util.Json(w, http.StatusCreated, util.Response{
			Data: hostel,
		})
	}
}

func (h *HttpHostelHandler) UpdateHostel(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Hostel
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
	errors, er2 := h.validate(r.Context(), &hostel)
	if er2 != nil {
		h.logError(r.Context(), er2.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		h.logError(r.Context(), er2.Error())
		util.Json(w, http.StatusUnprocessableEntity, errors)
		return
	}
	_, er3 := h.service.UpdateHostel(r.Context(), &hostel)
	if er3 != nil {
		h.logError(r.Context(), er3.Error())
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		}
	} else {
		util.Json(w, http.StatusOK, util.Response{
			Data: hostel,
		})
	}
}

func (h *HttpHostelHandler) DeleteHostel(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	res, err := h.service.DeleteHostel(r.Context(), code)
	if err != nil {
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
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

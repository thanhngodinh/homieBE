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
	"hostel-service/internal/util"
)

func NewHostelHandler(service service.HostelService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *HttpHostelHandler {
	return &HttpHostelHandler{service: service, validate: validate, logError: logError}
}

type HttpHostelHandler struct {
	service  service.HostelService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpHostelHandler) GetHostels(w http.ResponseWriter, r *http.Request) {
	pageIdxParam := r.URL.Query().Get("pageIdx")
	pageSizeParam := r.URL.Query().Get("pageSize")
	pageIdx, err := strconv.Atoi(pageIdxParam)
	if err != nil {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorWrongTypePageIdx.Error(),
		})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorWrongTypePageSize.Error(),
		})
		return
	}
	res, err := h.service.GetHostels(r.Context(), pageSize, pageIdx)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else {
		JSON(w, http.StatusOK, util.Response{
			Data:  res.Data,
			Total: res.Pagin.Total,
		})
	}
}

func (h *HttpHostelHandler) GetHostelById(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	hostel, err := h.service.GetHostelById(r.Context(), code)
	if err != nil {
		h.logError(r.Context(), err.Error())
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else if hostel == nil {
		JSON(w, http.StatusNotFound, util.Response{})
	} else {
		JSON(w, http.StatusOK, util.Response{
			Data: hostel,
		})
	}
}

func (h *HttpHostelHandler) CreateHostel(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Hostel
	er1 := json.NewDecoder(r.Body).Decode(&hostel)
	defer r.Body.Close()
	if er1 != nil {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: er1.Error(),
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
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	_, er3 := h.service.CreateHostel(r.Context(), &hostel)
	if er3 != nil {
		h.logError(r.Context(), er3.Error())
		if util.IsDefinedErrorType(er3) {
			JSON(w, http.StatusBadRequest, util.Response{
				Message: er3.Error(),
			})
		} else {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		}
	} else {
		JSON(w, http.StatusCreated, util.Response{
			Data: hostel,
		})
	}
}

func (h *HttpHostelHandler) UpdateHostel(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Hostel
	er1 := json.NewDecoder(r.Body).Decode(&hostel)
	defer r.Body.Close()
	if er1 != nil {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: er1.Error(),
		})
		return
	}
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	if len(hostel.Id) == 0 {
		hostel.Id = code
	} else if code != hostel.Id {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorCodeNotMatch.Error(),
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
		JSON(w, http.StatusUnprocessableEntity, errors)
		return
	}
	_, er3 := h.service.UpdateHostel(r.Context(), &hostel)
	if er3 != nil {
		h.logError(r.Context(), er3.Error())
		if util.IsDefinedErrorType(er3) {
			JSON(w, http.StatusBadRequest, util.Response{
				Message: er3.Error(),
			})
		} else {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		}
	} else {
		JSON(w, http.StatusOK, util.Response{
			Data: hostel,
		})
	}
}

func (h *HttpHostelHandler) DeleteHostel(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if len(code) == 0 {
		JSON(w, http.StatusBadRequest, util.Response{
			Message: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	res, err := h.service.DeleteHostel(r.Context(), code)
	if err != nil {
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
	} else {
		if res == 1 {
			JSON(w, http.StatusOK, util.Response{
				Data: fmt.Sprintf("delete %s successfully", code),
			})
		} else {
			JSON(w, http.StatusNotFound, util.Response{
				Data: fmt.Sprintf("not found %s", code),
			})
		}
	}
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

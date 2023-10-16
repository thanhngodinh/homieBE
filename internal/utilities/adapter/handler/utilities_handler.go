package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

	"hostel-service/internal/package/util"
	"hostel-service/internal/utilities/domain"
	"hostel-service/internal/utilities/service"
)

func NewUtilitiesHandler(
	service service.UtilitiesService,
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
	logError func(context.Context, string, ...map[string]interface{})) *HttpUtilitiesHandler {
	return &HttpUtilitiesHandler{
		service:  service,
		validate: validate,
		logError: logError}
}

type HttpUtilitiesHandler struct {
	service  service.UtilitiesService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func (h *HttpUtilitiesHandler) GetAllUtilities(w http.ResponseWriter, r *http.Request) {
	utilities, err := h.service.GetAllUtilities(r.Context())
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, utilities)
	}
}

func (h *HttpUtilitiesHandler) CreateUtilities(w http.ResponseWriter, r *http.Request) {
	var utilities domain.Utilities
	er1 := json.NewDecoder(r.Body).Decode(&utilities)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	errors, er2 := h.validate(r.Context(), &utilities)
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
	utilities.CreatedBy = r.Context().Value("userId").(string)
	_, er3 := h.service.CreateUtilities(r.Context(), &utilities)
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
			Data: utilities,
		})
	}
}

func (h *HttpUtilitiesHandler) UpdateUtilities(w http.ResponseWriter, r *http.Request) {
	var utilities domain.Utilities
	er1 := json.NewDecoder(r.Body).Decode(&utilities)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	code := mux.Vars(r)["id"]
	if len(code) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	if len(utilities.Id) == 0 {
		utilities.Id = code
	} else if code != utilities.Id {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeNotMatch.Error(),
		})
		return
	}
	errors, er2 := h.validate(r.Context(), &utilities)
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
	_, er3 := h.service.UpdateUtilities(r.Context(), &utilities)
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
			Data: utilities,
		})
	}
}

func (h *HttpUtilitiesHandler) DeleteUtilities(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["id"]
	if len(code) == 0 {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: util.ErrorCodeEmpty.Error(),
		})
		return
	}
	res, err := h.service.DeleteUtilities(r.Context(), code)
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

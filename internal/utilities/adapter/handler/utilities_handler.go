package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"hostel-service/internal/utilities/domain"
	"hostel-service/internal/utilities/service"
	"hostel-service/package/util"
)

func NewUtilitiesHandler(
	service service.UtilitiesService,
	validate *validator.Validate,
) *HttpUtilitiesHandler {
	return &HttpUtilitiesHandler{
		service:  service,
		validate: validate,
	}
}

type HttpUtilitiesHandler struct {
	service  service.UtilitiesService
	validate *validator.Validate
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
	utilities.CreatedBy = r.Context().Value("userId").(string)
	_, er3 := h.service.CreateUtilities(r.Context(), &utilities)
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
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
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
	_, er3 := h.service.UpdateUtilities(r.Context(), &utilities)
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
			Data: utilities,
		})
	}
}

func (h *HttpUtilitiesHandler) DeleteUtilities(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["id"]
	if len(code) == 0 {
		util.JsonBadRequest(w, util.ErrorCodeEmpty)
		return
	}
	res, err := h.service.DeleteUtilities(r.Context(), code)
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

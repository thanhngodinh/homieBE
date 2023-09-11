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

// GetHostels godoc
// @Summary      Get Hostels
// @Description  Return a list of the Hostels included the pagination
// @Tags         Hostels
// @Accept       json
// @Produce      json
// @Param        Authorization   header    string           true                   "The Authorization"        example(Bearer eyJhbGci...)
// @Param        pageIdx         query     string           true                   "The index of the page start from 0"           example(0)
// @Param        pageSize        query     string           true                   "The number of Hostels return on each page"  example(10)
// @Success      200             {object}  domain.GetHostelsResponse
// @Failure      500             {string}  string           "Internal Server Error"
// @Router       /hostels [get]
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
			Data: res,
			Code: http.StatusOK,
		})
	}
}

// GetHostel godoc
// @Summary      Get a Hostel
// @Description  Return a Hostel with the code
// @Tags         Hostels
// @Accept       json
// @Produce      json
// @Param        Authorization header    string           true                   "The Authorization"        example(Bearer eyJhbGc...)
// @Param        code          path      string           true                   "The code of the Hostel" example(07e7a76c-1bbb-11ed-861d-0242ac120002)
// @Success      200           {object}  domain.Hostel
// @Failure      500           {string}  string           "Internal Server Error"
// @Router       /hostels/{code} [get]
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

// CreateHostel godoc
// @Summary      Create a Hostel
// @Tags         Hostels
// @Accept       json
// @Produce      json
// @Param        Authorization   header    string           true                   "The Authorization"        example(Bearer eyJhbGci...)
// @Param        Hostel        body      domain.Hostel  true "Hostel to create"
// @Success      201             {object}  util.Response{value=int,data=domain.Hostel}
// @Failure      400             {string}  string          "Invalid character 's' looking for beginning of value"
// @Failure      500             {string}  string          "Internal Server Error"
// @Router       /hostels [post]
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

// UpdateHostel godoc
// @Summary      Update a Hostel
// @Tags         Hostels
// @Accept       json
// @Produce      json
// @Param        Authorization header    string          true             "The Authorization"        example(Bearer eyJhbGci...)
// @Param        code          path      string          true             "The code of the Hostel" example(07e7a76c-1bbb-11ed-861d-0242ac120002)
// @Param        Hostel      body      domain.Hostel true             "Hostel to update"
// @Success      200           {object}  util.Response{value=int,data=domain.Hostel}
// @Failure      400           {string}  string          "Invalid character 's' looking for beginning of value"
// @Failure      500           {string}  string          "Internal Server Error"
// @Router       /hostels/{code} [put]
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

// DeleteHostel godoc
// @Summary      Delete a Hostel
// @Tags         Hostels
// @Accept       json
// @Produce      json
// @Param        Authorization   header    string  true  "The Authorization"        example(Bearer eyJhbGci...)
// @Param        code            path      string  true  "The code of The Hostel"
// @Success      200             {string}  string  "1"
// @Failure      404             {string}  string  "0"
// @Failure      500             {string}  string  "Internal Server Error"
// @Router       /hostels/{code} [delete]
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

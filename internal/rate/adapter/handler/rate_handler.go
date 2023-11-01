package handler

import (
	"context"
	"encoding/json"
	"net/http"

	sv "github.com/core-go/core"
	"github.com/gorilla/mux"

	"hostel-service/internal/package/util"
	"hostel-service/internal/rate/domain"
	"hostel-service/internal/rate/service"
)

func NewRateHandler(
	service service.RateService,
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error),
) *HttpRateHandler {
	return &HttpRateHandler{
		service:  service,
		validate: validate,
	}
}

type HttpRateHandler struct {
	service  service.RateService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
}

func (h *HttpRateHandler) GetPostRate(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorIdEmpty)
		return
	}
	rate, err := h.service.GetPostRate(r.Context(), postId)
	if err != nil {
		util.JsonInternalError(w, err)
	} else {
		util.Json(w, http.StatusOK, rate)
	}
}

func (h *HttpRateHandler) CreateRate(w http.ResponseWriter, r *http.Request) {
	rate := &domain.Rate{}
	er1 := json.NewDecoder(r.Body).Decode(rate)
	defer r.Body.Close()
	if er1 != nil {
		util.Json(w, http.StatusBadRequest, util.Response{
			Status: er1.Error(),
		})
		return
	}
	errors, er2 := h.validate(r.Context(), rate)
	if er2 != nil {
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		util.Json(w, http.StatusUnprocessableEntity, errors)
		return
	}
	rate.UserId = r.Context().Value("userId").(string)
	_, er3 := h.service.CreateRate(r.Context(), rate)
	if er3 != nil {
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		}
	} else {
		util.Json(w, http.StatusCreated, util.Response{
			Data: rate,
		})
	}
}

func (h *HttpRateHandler) UpdateRate(w http.ResponseWriter, r *http.Request) {
	rate := &domain.Rate{}
	err := json.NewDecoder(r.Body).Decode(rate)
	defer r.Body.Close()
	if err != nil {
		util.JsonBadRequest(w, err)
		return
	}
	postId := mux.Vars(r)["postId"]
	if len(postId) == 0 {
		util.JsonBadRequest(w, util.ErrorIdEmpty)
		return
	} else if postId != rate.PostId {
		util.JsonBadRequest(w, util.ErrorCodeNotMatch)
		return
	}
	rate.UserId = r.Context().Value("userId").(string)
	errors, er2 := h.validate(r.Context(), rate)
	if er2 != nil {
		http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
		return
	}
	if len(errors) > 0 {
		util.Json(w, http.StatusUnprocessableEntity, errors)
		return
	}
	_, er3 := h.service.UpdateRate(r.Context(), rate)
	if er3 != nil {
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			http.Error(w, sv.InternalServerError, http.StatusInternalServerError)
			return
		}
	}
	util.JsonOK(w, rate)
}

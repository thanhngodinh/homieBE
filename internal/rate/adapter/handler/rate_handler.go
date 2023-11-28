package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"hostel-service/internal/rate/domain"
	"hostel-service/internal/rate/service"
	"hostel-service/pkg/util"
)

func NewRateHandler(
	service service.RateService,
	validate *validator.Validate,
) *HttpRateHandler {
	return &HttpRateHandler{
		service:  service,
		validate: validate,
	}
}

type HttpRateHandler struct {
	service  service.RateService
	validate *validator.Validate
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
		util.JsonOK(w, rate)
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
	rate.UserId = r.Context().Value("userId").(string)
	_, er3 := h.service.CreateRate(r.Context(), rate)
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
	_, er3 := h.service.UpdateRate(r.Context(), rate)
	if er3 != nil {
		if util.IsDefinedErrorType(er3) {
			util.Json(w, http.StatusBadRequest, util.Response{
				Status: er3.Error(),
			})
		} else {
			util.JsonInternalError(w, errors.New("internal server error"))
			return
		}
	}
	util.JsonOK(w, rate)
}

package util

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
}

const Success = "success"

func Json(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

func JsonOK(w http.ResponseWriter, data ...interface{}) error {
	if data != nil {
		if len(data) == 2 {
			return Json(w, http.StatusOK, Response{Data: data[0], Total: data[1].(int64)})
		}
		return Json(w, http.StatusOK, Response{Data: data[0]})
	}
	return Json(w, http.StatusOK, Response{Status: Success})
}

func JsonBadRequest(w http.ResponseWriter, err error, message ...string) error {
	logrus.Error(err)
	if message != nil {
		return Json(w, http.StatusBadRequest, Response{Status: err.Error(), Message: message[0]})
	}
	return Json(w, http.StatusBadRequest, Response{Status: err.Error(), Message: err.Error()})
}

func JsonInternalError(w http.ResponseWriter, err error, message ...string) error {
	logrus.Error(err)
	if message != nil {
		return Json(w, http.StatusInternalServerError, Response{Status: err.Error(), Message: message[0]})
	}
	return Json(w, http.StatusInternalServerError, Response{Status: err.Error(), Message: err.Error()})
}

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

const Success = "Success"
const InternalServerError = "Internal Server Error"

func Json(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

func JsonOK(w http.ResponseWriter, data ...interface{}) error {
	if data != nil {
		return Json(w, http.StatusOK, Response{Data: data[0]})
	}
	return Json(w, http.StatusOK, Response{Status: Success})
}

func JsonInternalError(w http.ResponseWriter, err error) error {
	logrus.Error(err)
	return Json(w, http.StatusInternalServerError, Response{Message: InternalServerError})
}

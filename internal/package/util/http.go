package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
}

func Json(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

const Success = "Success"

func JsonOK(w http.ResponseWriter) error {
	return Json(w, http.StatusOK, Response{Status: Success})
}

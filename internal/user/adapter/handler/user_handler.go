package handler

import (
	"context"
	"encoding/json"
	"net/http"

	sv "github.com/core-go/core"

	"hostel-service/internal/user/service"
)

func NewUserHandler(service service.UserService, validate func(context.Context, interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string, ...map[string]interface{})) *HttpUserHandler {
	return &HttpUserHandler{service: service, validate: validate, logError: logError}
}

type HttpUserHandler struct {
	service  service.UserService
	validate func(context.Context, interface{}) ([]sv.ErrorMessage, error)
	logError func(context.Context, string, ...map[string]interface{})
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

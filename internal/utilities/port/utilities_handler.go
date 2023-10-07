package port

import "net/http"

type UtilitiesHandler interface {
	GetAllUtilities(w http.ResponseWriter, r *http.Request)
	CreateUtilities(w http.ResponseWriter, r *http.Request)
	UpdateUtilities(w http.ResponseWriter, r *http.Request)
	DeleteUtilities(w http.ResponseWriter, r *http.Request)
}

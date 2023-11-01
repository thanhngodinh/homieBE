package port

import "net/http"

type RateHandler interface {
	GetPostRate(w http.ResponseWriter, r *http.Request)
	CreateRate(w http.ResponseWriter, r *http.Request)
	UpdateRate(w http.ResponseWriter, r *http.Request)
}

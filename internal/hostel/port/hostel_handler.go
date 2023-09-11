package port

import "net/http"

type HostelHandler interface {
	GetHostels(w http.ResponseWriter, r *http.Request)
	GetHostelById(w http.ResponseWriter, r *http.Request)
	CreateHostel(w http.ResponseWriter, r *http.Request)
	UpdateHostel(w http.ResponseWriter, r *http.Request)
	DeleteHostel(w http.ResponseWriter, r *http.Request)
}

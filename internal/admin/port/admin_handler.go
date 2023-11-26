package port

import "net/http"

type AdminHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetAdminInfo(w http.ResponseWriter, r *http.Request)
	ResetUserPassword(w http.ResponseWriter, r *http.Request)
	UpdateUserStatus(w http.ResponseWriter, r *http.Request)
	UpdatePostStatus(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

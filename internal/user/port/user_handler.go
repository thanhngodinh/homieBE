package port

import "net/http"

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetUserProfile(w http.ResponseWriter, r *http.Request)
	UpdateUserStatus(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	ResetUserPassword(w http.ResponseWriter, r *http.Request)
	SearchRoommates(w http.ResponseWriter, r *http.Request)
	GetRoommateById(w http.ResponseWriter, r *http.Request)
}

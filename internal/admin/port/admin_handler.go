package port

import "net/http"

type AdminHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	GetAdminProfile(w http.ResponseWriter, r *http.Request)

	SearchPosts(w http.ResponseWriter, r *http.Request)
	GetPostById(w http.ResponseWriter, r *http.Request)
	UpdatePostStatus(w http.ResponseWriter, r *http.Request)

	SearchUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUserStatus(w http.ResponseWriter, r *http.Request)
	ResetUserPassword(w http.ResponseWriter, r *http.Request)
}

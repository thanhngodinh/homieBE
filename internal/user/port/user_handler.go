package port

import "net/http"

type UserHandler interface {
	GetPostLikedByUser(w http.ResponseWriter, r *http.Request)
	UserLikePost(w http.ResponseWriter, r *http.Request)
}

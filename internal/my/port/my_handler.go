package port

import "net/http"

type MyHandler interface {
	GetMyPostLiked(w http.ResponseWriter, r *http.Request)
	GetMyPosts(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request)
}

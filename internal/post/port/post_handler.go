package port

import "net/http"

type PostHandler interface {
	GetSuggestPosts(w http.ResponseWriter, r *http.Request)
	SearchPosts(w http.ResponseWriter, r *http.Request)
	ElasticSearchPosts(w http.ResponseWriter, r *http.Request)
	GetPostById(w http.ResponseWriter, r *http.Request)
	GetCompare(w http.ResponseWriter, r *http.Request)
	CheckCreatePost(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

package app

import (
	"context"

	"github.com/gorilla/mux"

	internalMid "hostel-service/internal/middleware"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}

	userRouter := r.PathPrefix("/my").Subrouter()
	userRouter.HandleFunc("/liked-posts", app.My.GetMyPostLiked).Methods(GET)
	userRouter.HandleFunc("/posts", app.My.GetMyPosts).Methods(GET)
	userRouter.HandleFunc("/posts/{code}", app.Post.UpdatePost).Methods(PUT)
	userRouter.HandleFunc("/password", app.User.UpdatePassword).Methods(PUT)
	userRouter.HandleFunc("/profile", app.My.GetMyProfile).Methods(GET)
	userRouter.HandleFunc("/profile", app.My.UpdateMyProfile).Methods(PUT)
	userRouter.HandleFunc("/avatar", app.My.UpdateMyAvatar).Methods(PUT)
	userRouter.Use(internalMid.Authenticate)

	hostelRouter := r.PathPrefix("/posts").Subrouter()
	hostelRouter.HandleFunc("", app.Post.CreatePost).Methods(POST)
	hostelRouter.HandleFunc("/like/{postId}", app.My.LikePost).Methods(POST)
	hostelRouter.HandleFunc("/{code}", app.Post.UpdatePost).Methods(PUT)
	hostelRouter.HandleFunc("/{code}", app.Post.DeletePost).Methods(DELETE)
	hostelRouter.Use(internalMid.Authenticate)

	hostelPublicRouter := r.PathPrefix("/posts").Subrouter()
	hostelPublicRouter.HandleFunc("", app.Post.GetPosts).Methods(GET)
	hostelPublicRouter.HandleFunc("/search", app.Post.SearchPosts).Methods(POST)
	hostelPublicRouter.HandleFunc("/suggest", app.Post.GetSuggestPosts).Methods(GET)
	hostelPublicRouter.HandleFunc("/{code}", app.Post.GetPostById).Methods(GET)
	hostelPublicRouter.Use(internalMid.PublicAuth)

	roommateRouter := r.PathPrefix("/roommates").Subrouter()
	roommateRouter.HandleFunc("/search", app.User.SearchRoommates).Methods(POST)
	roommateRouter.HandleFunc("/{userId}", app.User.GetRoommateById).Methods(GET)
	// roommateRouter.Use(internalMid.PublicAuth)

	r.HandleFunc("/utilities", app.Utilities.GetAllUtilities).Methods(GET)
	utilitiesRouter := r.PathPrefix("/utilities").Subrouter()
	utilitiesRouter.HandleFunc("", app.Utilities.CreateUtilities).Methods(POST)
	utilitiesRouter.HandleFunc("/{id}", app.Utilities.UpdateUtilities).Methods(PUT)
	utilitiesRouter.HandleFunc("/{id}", app.Utilities.DeleteUtilities).Methods(DELETE)
	utilitiesRouter.Use(internalMid.Authenticate)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", app.User.Register).Methods(POST)
	authRouter.HandleFunc("/login", app.User.Login).Methods(POST)

	chatRouter := r.PathPrefix("/chat").Subrouter()
	chatRouter.HandleFunc("", app.Chat.InitConversation)

	return nil
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

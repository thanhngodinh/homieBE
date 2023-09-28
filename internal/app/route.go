package app

import (
	"context"

	"github.com/gorilla/mux"

	internalMid "hostel-service/internal/middleware"
	// httpSwagger "github.com/swaggo/http-swagger"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}

	userRouter := r.PathPrefix("/my").Subrouter()
	userRouter.HandleFunc("/liked-posts", app.My.GetMyPostLiked).Methods(GET)
	userRouter.HandleFunc("/posts", app.My.GetMyPosts).Methods(GET)
	userRouter.HandleFunc("/like/{postId}", app.My.LikePost).Methods(POST)
	userRouter.Use(internalMid.Authenticate)

	r.HandleFunc("/hostels", app.Hostel.GetHostels).Methods(GET)

	hostelRouter := r.PathPrefix("/hostels").Subrouter()
	hostelRouter.HandleFunc("/suggest", app.Hostel.GetSuggestHostels).Methods(GET)
	hostelRouter.HandleFunc("", app.Hostel.CreateHostel).Methods(POST)
	hostelRouter.HandleFunc("/{code}", app.Hostel.UpdateHostel).Methods(PUT)
	// hostelRouter.HandleFunc("/{code}", app.Hostel.DeleteHostel).Methods(DELETE)
	hostelRouter.Use(internalMid.Authenticate)

	hostelPublicRouter := r.PathPrefix("/hostels").Subrouter()
	hostelPublicRouter.HandleFunc("/search", app.Hostel.SearchHostels).Methods(GET)
	hostelPublicRouter.HandleFunc("/{code}", app.Hostel.GetHostelById).Methods(GET)
	hostelPublicRouter.Use(internalMid.PublicAuth)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", app.User.Register).Methods(POST)
	authRouter.HandleFunc("/login", app.User.Login).Methods(POST)

	// r.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	return nil
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

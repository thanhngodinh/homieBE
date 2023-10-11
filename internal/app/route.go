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
	userRouter.HandleFunc("/posts/{code}", app.Hostel.UpdateHostel).Methods(PUT)
	userRouter.HandleFunc("/like/{postId}", app.My.LikePost).Methods(POST)
	userRouter.HandleFunc("/password", app.User.UpdatePassword).Methods(PUT)
	userRouter.Use(internalMid.Authenticate)

	hostelRouter := r.PathPrefix("/hostels").Subrouter()
	hostelRouter.HandleFunc("", app.Hostel.CreateHostel).Methods(POST)
	hostelRouter.HandleFunc("/{code}", app.Hostel.UpdateHostel).Methods(PUT)
	hostelRouter.HandleFunc("/{code}", app.Hostel.DeleteHostel).Methods(DELETE)
	hostelRouter.Use(internalMid.Authenticate)

	hostelPublicRouter := r.PathPrefix("/hostels").Subrouter()
	hostelPublicRouter.HandleFunc("", app.Hostel.GetHostels).Methods(GET)
	hostelPublicRouter.HandleFunc("/search", app.Hostel.SearchHostels).Methods(POST)
	hostelPublicRouter.HandleFunc("/suggest", app.Hostel.GetSuggestHostels).Methods(GET)
	hostelPublicRouter.HandleFunc("/{code}", app.Hostel.GetHostelById).Methods(GET)
	hostelPublicRouter.Use(internalMid.PublicAuth)

	r.HandleFunc("/utilities", app.Utilities.GetAllUtilities).Methods(GET)
	utilitiesRouter := r.PathPrefix("/utilities").Subrouter()
	utilitiesRouter.HandleFunc("", app.Utilities.CreateUtilities).Methods(POST)
	utilitiesRouter.HandleFunc("/{id}", app.Utilities.UpdateUtilities).Methods(PUT)
	utilitiesRouter.HandleFunc("/{id}", app.Utilities.DeleteUtilities).Methods(DELETE)
	utilitiesRouter.Use(internalMid.Authenticate)

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

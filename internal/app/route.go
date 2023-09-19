package app

import (
	"context"

	. "github.com/core-go/core"
	"github.com/gorilla/mux"

	internalMid "hostel-service/internal/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.HandleFunc("/health", app.Health.Check).Methods(GET)

	hostelSuggestRouter := r.PathPrefix("post/suggest").Subrouter()
	hostelSuggestRouter.HandleFunc("", app.Hostel.GetHostels).Methods(GET)
	hostelSuggestRouter.Use(internalMid.Authenticate)

	hostelRouter := r.PathPrefix("/hostels").Subrouter()
	hostelRouter.HandleFunc("", app.Hostel.GetHostels).Methods(GET)
	hostelRouter.HandleFunc("/{code}", app.Hostel.GetHostelById).Methods(GET)
	hostelRouter.HandleFunc("", app.Hostel.CreateHostel).Methods(POST)
	hostelRouter.HandleFunc("/{code}", app.Hostel.UpdateHostel).Methods(PUT)
	hostelRouter.HandleFunc("/{code}", app.Hostel.DeleteHostel).Methods(DELETE)
	// hostelRouter.Use(internalMid.Authenticate)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", app.Auth.Register).Methods(POST)
	authRouter.HandleFunc("/login", app.Auth.Login).Methods(POST)

	r.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	return nil
}

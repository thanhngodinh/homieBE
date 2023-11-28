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

	adminUserRouter := r.PathPrefix("/users").Subrouter()
	adminUserRouter.HandleFunc("/{userId}/reset-password", app.User.ResetUserPassword).Methods(PATCH)
	adminUserRouter.HandleFunc("/{userId}/disable", app.User.UpdateUserStatus).Methods(PATCH)
	adminUserRouter.HandleFunc("/{userId}/active", app.User.UpdateUserStatus).Methods(PATCH)
	adminUserRouter.Use(internalMid.AdminAuthenticate)

	adminPostRouter := r.PathPrefix("/posts").Subrouter()
	adminPostRouter.HandleFunc("/{postId}/verify", app.Post.UpdatePostStatus).Methods(PATCH)
	adminPostRouter.HandleFunc("/{postId}/active", app.Post.UpdatePostStatus).Methods(PATCH)
	adminPostRouter.HandleFunc("/{postId}/disable", app.Post.UpdatePostStatus).Methods(PATCH)
	adminPostRouter.Use(internalMid.AdminAuthenticate)

	myRouter := r.PathPrefix("/my").Subrouter()
	myRouter.HandleFunc("/liked-posts", app.My.GetMyPostLiked).Methods(GET)
	myRouter.HandleFunc("/posts", app.My.GetMyPosts).Methods(GET)
	myRouter.HandleFunc("/posts/{code}", app.Post.UpdatePost).Methods(PUT)
	myRouter.HandleFunc("/password", app.User.UpdatePassword).Methods(PATCH)
	myRouter.HandleFunc("/verify-phone", app.User.VerifyPhone).Methods(PATCH)
	myRouter.HandleFunc("/verify-otp", app.User.VerifyPhoneOTP).Methods(PATCH)
	myRouter.HandleFunc("/password", app.User.UpdatePassword).Methods(PATCH)
	myRouter.HandleFunc("/profile", app.My.GetMyProfile).Methods(GET)
	myRouter.HandleFunc("/profile", app.My.UpdateMyProfile).Methods(PUT)
	myRouter.HandleFunc("/avatar", app.My.UpdateMyAvatar).Methods(PUT)
	myRouter.Use(internalMid.Authenticate)

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{userId}", app.User.GetUserProfile).Methods(GET)
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
	hostelPublicRouter.HandleFunc("/esearch", app.Post.ElasticSearchPosts).Methods(POST)
	hostelPublicRouter.HandleFunc("/suggest", app.Post.GetSuggestPosts).Methods(GET)
	hostelPublicRouter.HandleFunc("/compare/{post1}/{post2}", app.Post.GetCompare).Methods(GET)
	hostelPublicRouter.HandleFunc("/{code}", app.Post.GetPostById).Methods(GET)
	hostelPublicRouter.HandleFunc("/check-create", app.Post.CheckCreatePost).Methods(POST)
	hostelPublicRouter.Use(internalMid.PublicAuth)

	rateRouter := r.PathPrefix("/rates").Subrouter()
	rateRouter.HandleFunc("", app.Rate.CreateRate).Methods(POST)
	rateRouter.HandleFunc("/{postId}", app.Rate.UpdateRate).Methods(PUT)
	rateRouter.Use(internalMid.Authenticate)

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

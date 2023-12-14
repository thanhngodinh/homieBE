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

	//Admin
	adminRouter := r.PathPrefix("/admin").Subrouter()
	r.HandleFunc("/admin/login", app.Admin.Login).Methods(POST)

	adminRouter.HandleFunc("/profile", app.Admin.GetAdminProfile).Methods(GET)
	adminRouter.HandleFunc("/password", app.Admin.UpdatePassword).Methods(PATCH)
	adminRouter.Use(internalMid.AdminAuthenticate)

	adminUserRouter := adminRouter.PathPrefix("/users").Subrouter()
	adminUserRouter.HandleFunc("/search", app.Admin.SearchUsers).Methods(POST)
	adminUserRouter.HandleFunc("/{userId}", app.Admin.GetUserById).Methods(GET)
	adminUserRouter.HandleFunc("/{userId}/reset-password", app.Admin.ResetUserPassword).Methods(PATCH)
	adminUserRouter.HandleFunc("/{userId}/disable", app.Admin.UpdateUserStatus).Methods(PATCH)
	adminUserRouter.HandleFunc("/{userId}/active", app.Admin.UpdateUserStatus).Methods(PATCH)
	// adminUserRouter.Use(internalMid.AdminAuthenticate)

	adminPostRouter := adminRouter.PathPrefix("/posts").Subrouter()
	adminPostRouter.HandleFunc("/search", app.Admin.SearchPosts).Methods(POST)
	adminPostRouter.HandleFunc("/{postId}", app.Admin.GetPostById).Methods(GET)
	adminPostRouter.HandleFunc("/{postId}/verify", app.Admin.UpdatePostStatus).Methods(PATCH)
	adminPostRouter.HandleFunc("/{postId}/active", app.Admin.UpdatePostStatus).Methods(PATCH)
	adminPostRouter.HandleFunc("/{postId}/disable", app.Admin.UpdatePostStatus).Methods(PATCH)
	// adminPostRouter.Use(internalMid.AdminAuthenticate)

	utilitiesRouter := r.PathPrefix("/utilities").Subrouter()
	utilitiesRouter.HandleFunc("", app.Utilities.CreateUtilities).Methods(POST)
	utilitiesRouter.HandleFunc("/{id}", app.Utilities.UpdateUtilities).Methods(PUT)
	utilitiesRouter.HandleFunc("/{id}", app.Utilities.DeleteUtilities).Methods(DELETE)
	utilitiesRouter.Use(internalMid.AdminAuthenticate)

	//User
	myRouter := r.PathPrefix("/my").Subrouter()
	myRouter.HandleFunc("/liked-posts", app.My.GetMyPostLiked).Methods(GET)
	myRouter.HandleFunc("/posts", app.My.GetMyPosts).Methods(GET)
	myRouter.HandleFunc("/posts/{code}", app.Post.UpdatePost).Methods(PUT)
	myRouter.HandleFunc("/verify-phone", app.User.VerifyPhone).Methods(PATCH)
	myRouter.HandleFunc("/verify-otp", app.User.VerifyPhoneOTP).Methods(PATCH)
	myRouter.HandleFunc("/password", app.User.UpdatePassword).Methods(PATCH)
	myRouter.HandleFunc("/profile", app.My.GetMyProfile).Methods(GET)
	myRouter.HandleFunc("/profile", app.My.UpdateMyProfile).Methods(PUT)
	myRouter.HandleFunc("/avatar", app.My.UpdateMyAvatar).Methods(PATCH)
	myRouter.Use(internalMid.Authenticate)

	hostelRouter := r.PathPrefix("/posts").Subrouter()
	hostelRouter.HandleFunc("", app.Post.CreatePost).Methods(POST)
	hostelRouter.HandleFunc("/like/{postId}", app.My.LikePost).Methods(POST)
	hostelRouter.HandleFunc("/{code}", app.Post.UpdatePost).Methods(PUT)
	hostelRouter.HandleFunc("/{code}", app.Post.DeletePost).Methods(DELETE)
	hostelRouter.Use(internalMid.Authenticate)

	hostelPublicRouter := r.PathPrefix("/posts").Subrouter()
	hostelPublicRouter.HandleFunc("/search", app.Post.SearchPosts).Methods(POST)
	hostelPublicRouter.HandleFunc("/esearch", app.Post.ElasticSearchPosts).Methods(POST)
	hostelPublicRouter.HandleFunc("/suggest", app.Post.GetSuggestPosts).Methods(GET)
	hostelPublicRouter.HandleFunc("/compare/{post1}/{post2}", app.Post.GetCompare).Methods(GET)
	hostelPublicRouter.HandleFunc("/{code}", app.Post.GetPostById).Methods(GET)
	hostelPublicRouter.HandleFunc("/check-create", app.Post.CheckCreatePost).Methods(POST)
	hostelPublicRouter.Use(internalMid.PublicAuth)

	rateRouter := r.PathPrefix("/rates").Subrouter()
	rateRouter.HandleFunc("/{postId}", app.Rate.CreateRate).Methods(POST)
	rateRouter.HandleFunc("/{postId}", app.Rate.UpdateRate).Methods(PATCH)
	rateRouter.Use(internalMid.Authenticate)

	roommateRouter := r.PathPrefix("/roommates").Subrouter()
	roommateRouter.HandleFunc("/search", app.User.SearchRoommates).Methods(POST)
	roommateRouter.HandleFunc("/{userId}", app.User.GetRoommateById).Methods(GET)
	roommateRouter.Use(internalMid.PublicAuth)

	r.HandleFunc("/utilities", app.Utilities.GetAllUtilities).Methods(GET)

	r.HandleFunc("/users/{userId}", app.User.GetUserProfile).Methods(GET)

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

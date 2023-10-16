package main

import (
	"context"
	"fmt"
	_ "hostel-service/docs"
	"net/http"

	"github.com/core-go/config"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	// "github.com/core-go/core/cors"
	"github.com/core-go/log"
	"github.com/gorilla/mux"

	"hostel-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	log.Initialize(conf.Log)

	err = app.Route(r, context.Background(), conf)
	if err != nil {
		panic(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost", "http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"*"},
		// AllowCredentials: true,
	})
	handler := c.Handler(r)


	logrus.Infof("Server start at port: %v", *conf.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", *conf.Server.Port), handler)
}

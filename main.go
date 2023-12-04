package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/core-go/config"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"hostel-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./configs")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }

	r := mux.NewRouter()

	err = app.Route(r, context.Background(), conf)
	if err != nil {
		panic(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*", "http://localhost", "http://localhost:3000", "http://localhost:3030"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"*"},
		// AllowCredentials: true,
	})
	handler := c.Handler(r)

	logrus.Infof("Server start at port: %v", *conf.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", *conf.Server.Port), handler)
}

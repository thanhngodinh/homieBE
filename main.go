package main

import (
	"context"
	"fmt"
	_ "hostel-service/docs"

	"github.com/core-go/config"
	"github.com/core-go/core"
	"github.com/core-go/core/cors"
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

	c := cors.New(conf.Allow)
	handler := c.Handler(r)
	fmt.Println(core.ServerInfo(conf.Server))
	server := core.CreateServer(conf.Server, handler)

	if err = server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}

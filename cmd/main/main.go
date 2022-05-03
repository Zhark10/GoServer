package main

import (
	"github.com/julienschmidt/httprouter"
	"goServer/internal/users"
	"goServer/pkg/logging"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()
	logger.Info("register user handle")
	userHandler := users.NewHandler(logger)
	userHandler.Register(router)
	startServer(router)
}

func startServer(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start app")
	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening on port 8888")
	logger.Fatal(server.Serve(listener))
}

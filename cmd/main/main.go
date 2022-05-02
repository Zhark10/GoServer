package main

import (
	"github.com/julienschmidt/httprouter"
	"goServer/internal/users"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	log.Println("register user handle")
	userHandler := users.NewHandler()
	userHandler.Register(router)
	startServer(router)
}

func startServer(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}

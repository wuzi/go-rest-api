package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	router *chi.Mux
}

func run() error {
	server := newServer()
	fmt.Println("Server started at http://localhost:3333")
	http.ListenAndServe(":3333", server.router)
	return nil
}

func newServer() *server {
	s := &server{}
	s.routes()
	return s
}

package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	router *chi.Mux
}

func run() error {
	flag.Parse()
	server := newServer()
	http.ListenAndServe(":3333", server.router)
	return nil
}

func newServer() *server {
	s := &server{}
	s.routes()
	return s
}

package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func (s *server) routes() {
	s.router = chi.NewRouter()

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.URLFormat)
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/polls", func(rt chi.Router) {
		rt.Get("/", ListPolls)
		rt.Post("/", CreatePoll)
		rt.Route("/{id:[0-9]+}", func(r chi.Router) {
			r.Get("/", SinglePoll)
			r.Put("/", UpdatePoll)
			r.Patch("/", UpdatePoll)
			r.Delete("/", DeletePoll)
		})
	})

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("EasyPoll API Online"))
	})
}

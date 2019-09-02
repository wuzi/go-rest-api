package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

func (s *server) routes() {
	s.router = chi.NewRouter()

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.URLFormat)
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	token := jwtauth.New("HS256", []byte("a2CmyEmA0m"), nil)
	s.router.Use(jwtauth.Verifier(token))

	s.router.Route("/polls", func(r chi.Router) {
		r.Get("/", ListPolls)

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Authenticator)
			r.Post("/", CreatePoll)
		})

		r.Route("/{id:[0-9]+}", func(r chi.Router) {
			r.Get("/", SinglePoll)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Authenticator)
				r.Put("/", UpdatePoll)
				r.Patch("/", UpdatePoll)
				r.Delete("/", DeletePoll)
			})
		})
	})

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("EasyPoll API Online"))
	})
}

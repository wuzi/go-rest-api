package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Poll is the struct of the poll
type Poll struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

// PollResponse is the struct of the response of polls
type PollResponse struct {
	*Poll
}

// polls is a slice with the registered polls
// it is used to simulate a database. It should
// be changed to database later on
var polls = []*Poll{
	{ID: 1, Title: "Project Meeting", Slug: "project-meeting"},
	{ID: 2, Title: "Favorite Activities", Slug: "favotire-activities"},
	{ID: 3, Title: "Favorite Food", Slug: "favorite-food"},
}

// PollRouter mount the routes used for the poll resource
func PollRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", ListPolls)
	return r
}

// ListPolls appends multiple PollResponse into a list and render them as JSON
func ListPolls(w http.ResponseWriter, r *http.Request) {
	list := []render.Renderer{}
	for _, poll := range polls {
		list = append(list, &PollResponse{Poll: poll})
	}

	if err := render.RenderList(w, r, list); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// Render do a pre-processing before a response is marshalled and sent across the wire
func (rd *PollResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package main

import (
	"errors"
	"math/rand"
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

// PollRequest is the struct of the request of polls
type PollRequest struct {
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
	r.Post("/", CreatePoll)
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

// CreatePoll appends a new poll to the polls list
func CreatePoll(w http.ResponseWriter, r *http.Request) {
	data := &PollRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	poll := data.Poll
	dbNewPoll(poll)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &PollResponse{Poll: poll})
}

// Render do a pre-processing before a response is marshalled and sent across the wire
func (rd *PollResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind validates the poll request
func (rd *PollRequest) Bind(r *http.Request) error {
	if rd.Poll == nil {
		return errors.New("missing fields")
	}
	return nil
}

// Repository methods
func dbNewPoll(poll *Poll) (int64, error) {
	poll.ID = rand.Int63n(100) + 10
	polls = append(polls, poll)
	return poll.ID, nil
}

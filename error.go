package main

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse struct data
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status,omitempty"` // user-level status message
	AppCode    int64  `json:"code,omitempty"`   // application-specific error code
	ErrorText  string `json:"error,omitempty"`  // application-level error message, for debugging
}

// Render sets the status of the http code
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest returns an error for invalid requests
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrRender returns an error response for error when rendering response
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound returns a 404 error saying "Resource not found."
var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

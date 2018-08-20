package handlers

import (
	"net/http"

	"github.com/go-chi/render"

	api "github.com/arkhaix/lit-reader/api/scraper"
	"time"
)

var (
	// ScraperClient is the gRPC client for communicating with the scraper service.
	// Set this before using the handlers
	ScraperClient api.ScraperClient

	// ScraperTimeout is the gRPC timeout.
	// Set this before using the handlers
	ScraperTimeout time.Duration
)

type errResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *errResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func errInternalError(err error) render.Renderer {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal error",
		ErrorText:      err.Error(),
	}
}

func errInvalidRequest(err error) render.Renderer {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func errNotFound() render.Renderer {
	return &errResponse{
		HTTPStatusCode: 404,
		StatusText:     "Resource not found.",
	}
}

func errRender(err error) render.Renderer {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

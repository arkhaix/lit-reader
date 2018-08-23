package common

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse is the general response struct used by handlers
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render renders an ErrResponse
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInternalError returns an ErrResponse with status 500
func ErrInternalError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal error",
		ErrorText:      err.Error(),
	}
}

// ErrInvalidRequest returns an ErrResponse with status 400
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound returns an ErrResponse with status 404
func ErrNotFound() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: 404,
		StatusText:     "Resource not found.",
	}
}

// ErrRender returns an ErrResponse with status 422
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

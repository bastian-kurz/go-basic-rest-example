package internal

import (
	"github.com/go-chi/render"
	"net/http"
)

type ResponseError struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrorInvalidRequest(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid Request",
		ErrorText:      err.Error(),
	}
}

func ErrorResourceAlreadyExists(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		StatusText:     "Resource already exists",
		ErrorText:      err.Error(),
	}
}

func ErrorInternalServerError(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Something wentg wrong",
		ErrorText:      err.Error(),
	}
}

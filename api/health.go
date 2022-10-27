package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Health(r chi.Router) {
	r.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		// 🤪
	})
}

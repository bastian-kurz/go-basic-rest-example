package api

import (
	"context"
	"errors"
	"github.com/bastian-kurz/basic-rest-example/internal"
	"github.com/bastian-kurz/basic-rest-example/internal/entity/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func User(r chi.Router) {
	repository := user.NewRepository()
	service := user.NewService(repository)

	r.Group(func(r chi.Router) {
		r.Route("/api/user", func(r chi.Router) {
			// List of users
			r.Get("/", service.List)

			// Create a new user
			r.Post("/", service.Create)

			r.Route("/{userId}", func(r chi.Router) {
				r.Use(articleCtx)
				r.Get("/", service.Get)
			})
		})
	})
}

// to prevent should not use built-in type string as key for value; define your own type to avoid collisions (SA1029)
// staticcheck message
type userIdCtxKey string

func articleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		userId := chi.URLParam(r, "userId")
		if userId == "" {
			err = errors.New("missing required userId parameter")
			_ = render.Render(w, r, internal.ErrorInvalidRequest(err))
			return
		}
		var key userIdCtxKey = "userId"
		ctx := context.WithValue(r.Context(), key, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

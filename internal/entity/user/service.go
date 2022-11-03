package user

import (
	"github.com/bastian-kurz/basic-rest-example/internal"
	"github.com/bastian-kurz/basic-rest-example/internal/logger"
	"github.com/go-chi/render"
	"net/http"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	data := &Entity{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, internal.ErrorInvalidRequest(err))
		return
	}

	user := data
	user, err := s.repository.Create(r.Context(), user)
	if err != nil {
		logger.Log().Error(err.Error())
		_ = render.Render(w, r, internal.ErrorInternalServerError(err))
		return
	}

	res := Response{Data: []Items{{user}}}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, &res)
}

func (s *Service) List(w http.ResponseWriter, r *http.Request) {
	users, err := s.repository.List(r.Context())
	if err != nil {
		logger.Log().Error(err.Error())
		_ = render.Render(w, r, internal.ErrorInternalServerError(err))
	}

	res := Response{Data: users}

	render.Status(r, http.StatusOK)
	_ = render.Render(w, r, &res)
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	user, err := s.repository.Get(r.Context())
	if err != nil {
		logger.Log().Error(err.Error())
		_ = render.Render(w, r, internal.ErrorInternalServerError(err))
	}

	res := Response{Data: []Items{{user}}}
	render.Status(r, http.StatusOK)
	_ = render.Render(w, r, &res)
}

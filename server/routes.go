package server

import (
	"github.com/bastian-kurz/basic-rest-example/api"
)

func (s *Server) SetupRoutes() {
	api.Health(s.mux)
}

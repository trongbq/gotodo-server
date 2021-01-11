package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) endpointNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	s.respond(w, r, http.StatusNotFound, newErrResp(ErrCodeNotFound, "Endpoint not found"))
}

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Health string `json:"health"`
	}
	s.respond(w, r, http.StatusOK, response{"OK"})
}

// install method handles routes registration and also middleware setup
func (s Server) install() {
	s.router.NotFound(http.HandlerFunc(s.endpointNotFoundHandler))
	// Used by monitoring service to check health of running server
	s.router.Get("/monitor/check", http.HandlerFunc(s.healthCheckHandler))

	s.router.Route("/api/users", func(r chi.Router) {
		r.Get("/current", s.getCurrentUser)
		r.Post("/", s.registerUser)
	})
}

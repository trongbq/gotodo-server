package api

import (
	"net/http"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal/api/middleware"
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
	s.router.Use(chimiddleware.RequestID)
	s.router.NotFound(http.HandlerFunc(s.endpointNotFoundHandler))
	// Used by monitoring service to check health of running server
	s.router.Get("/monitor/check", http.HandlerFunc(s.healthCheckHandler))

	// Public endpoints
	s.router.Post("/api/users", s.registerUser)
	s.router.Post("/api/signin", s.signIn)

	// Secured endpoints
	s.router.Group(func(r chi.Router) {
		r.Use(middleware.Auth(s.auth))
		r.Get("/api/users/current", s.getCurrentUser)
		r.Route("/api/todos", func(r chi.Router) {
			r.Get("/", s.getTodoList)
			r.Post("/", s.addTodo)
			r.Put("/{todoID:[0-9+]}", s.updateTodo)
			r.Put("/{todoID:[0-9+]}/complete", s.completeTodo)
			r.Delete("/{todoID:[0-9+]}", s.deleteTodo)
		})
	})

	s.routerWalk()
}

func (s *Server) routerWalk() {
	chi.Walk(s.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.WithFields(log.Fields{
			"method": method,
			"route":  route,
		}).Debug()
		return nil
	})
}

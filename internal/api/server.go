package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/trongbq/gotodo-server/internal/api/auth"
	"github.com/trongbq/gotodo-server/internal/database"
)

// ServerConfig contains configuration values needed for Server instance running
type ServerConfig struct {
	Env          string
	AuthTokenKey string
}

type Server struct {
	conf   ServerConfig
	router *chi.Mux
	db     *database.DB
	auth   *auth.TokenIssuer
}

func NewServer(conf ServerConfig, db *database.DB) *Server {
	router := chi.NewRouter()
	auth := auth.NewTokenIssuer([]byte(conf.AuthTokenKey), "GoTodoServerIssuer")

	s := Server{
		conf:   conf,
		router: router,
		db:     db,
		auth:   auth,
	}
	// Install config all routes in the api server
	s.install()

	return &s
}

// ServeHTTP is just a wrapper for router but it makes Server become normal Handler
// It hides underlying logic of whichever router that Server are using
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

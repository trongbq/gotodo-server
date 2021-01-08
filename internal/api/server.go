package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trongbq/gotodo-server/internal/config"
	"github.com/trongbq/gotodo-server/internal/database"
)

// ServerConfig contains configuration values needed for Server instance running
type ServerConfig struct {
	Env string
}

type Server struct {
	conf   ServerConfig
	router *gin.Engine
	db     *database.DB
}

func NewServer(conf ServerConfig, db *database.DB) *Server {
	// Set running mode for gin depends on server's env
	if conf.Env == config.BetaEnv || conf.Env == config.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	} else if conf.Env == config.TestEnv {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()

	s := Server{
		conf:   conf,
		router: router,
		db:     db,
	}
	// Install config all routes in the api server
	s.install()

	return &s
}

func (s Server) ListenAndServe() error {
	return s.router.Run()
}

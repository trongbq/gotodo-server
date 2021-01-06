package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trongbq/gotodo-server/internal/config"
)

type Server struct {
	conf   ServerConfig
	router *gin.Engine
}

func NewServer(conf ServerConfig) *Server {
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
	}
	// Install config all routes in the api server
	s.install()

	return &s
}

func (s Server) ListenAndServe() error {
	return s.router.Run()
}

package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	conf   ServerConfig
	router *gin.Engine
}

func NewServer(conf ServerConfig) *Server {
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

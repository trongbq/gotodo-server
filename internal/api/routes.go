package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	endpointNotFoundHandler = func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:      ErrorCodeNotFound,
			Message:   "Endpoint not found",
			Timestamp: time.Now(),
		})
	}

	healthCheckHandler = func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "OK",
		})
	}
)

func (s Server) install() {
	s.router.NoRoute(endpointNotFoundHandler)
	// Used by monitoring service to check health of running server
	s.router.GET("/monitor/check", healthCheckHandler)
	s.router.POST("/api/users", s.RegisterUser)
	s.router.GET("/api/users/current", s.GetCurrentUser)
}

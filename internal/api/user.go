package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trongbq/gotodo-server/internal"
	"github.com/trongbq/gotodo-server/internal/database"
)

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (s *Server) GetCurrentUser(c *gin.Context) {
	user, err := s.db.GetUser(c, 1)
	if err == database.ErrNoRecordFound {
		c.JSON(http.StatusNotFound, newErrorResponse(ErrorCodeNotFound, "Current user not found"))
		return
	}
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrorResponse(ErrorCodeInternalError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) RegisterUser(c *gin.Context) {
	var req UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, newErrorResponse(ErrorCodeBadRequest, err.Error()))
		return
	}
	user := internal.User{
		Email:     req.Email,
		Username:  req.Username,
		CreatedAt: time.Now(),
	}
	id, err := s.db.InsertUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = id
	c.JSON(http.StatusOK, user)
}

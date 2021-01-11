package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/trongbq/gotodo-server/internal"
	"github.com/trongbq/gotodo-server/internal/database"
)

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (s *Server) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.db.GetUser(r.Context(), 1)
	if err == database.ErrNoRecordFound {
		s.respond(w, r, http.StatusNotFound, newErrResp(ErrCodeNotFound, "Current user not found"))
		return
	}
	fmt.Println(err)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	s.respond(w, r, http.StatusOK, user)
}

func (s *Server) registerUser(w http.ResponseWriter, r *http.Request) {
	var req UserRegisterRequest
	if err := s.decode(w, r, &req); err != nil {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, err.Error()))
		return
	}
	user := internal.User{
		Email:     req.Email,
		Username:  req.Username,
		CreatedAt: time.Now(),
	}
	id, err := s.db.InsertUser(r.Context(), user)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	user.ID = id
	s.respond(w, r, http.StatusOK, user)
}

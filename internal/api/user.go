package api

import (
	"net/http"
	"time"

	"github.com/trongbq/gotodo-server/internal"
	"github.com/trongbq/gotodo-server/internal/api/auth"
	"github.com/trongbq/gotodo-server/internal/api/request"
	"github.com/trongbq/gotodo-server/internal/database"
)

type userRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	log, _ := request.LogFrom(r.Context())
	user, _ := request.UserFrom(r.Context())
	log.Debug("Get current user with id: %v", user.ID)
	s.respond(w, r, http.StatusOK, user)
}

func (s *Server) registerUser(w http.ResponseWriter, r *http.Request) {
	log, _ := request.LogFrom(r.Context())
	var req userRegisterRequest
	if err := s.decode(w, r, &req); err != nil {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, err.Error()))
		return
	}
	passwd, err := auth.HashPassword(req.Password)
	if err != nil {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, err.Error()))
		return
	}
	user := database.DBUser{
		User: internal.User{
			Email:     req.Email,
			Username:  req.Username,
			CreatedAt: time.Now(),
		},
		Password: passwd,
	}
	id, err := s.db.InsertUser(r.Context(), user)
	if err != nil {
		log.Errorf("Register user failed due to %s", err.Error())
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	user.ID = id
	s.respond(w, r, http.StatusOK, user.User)
}

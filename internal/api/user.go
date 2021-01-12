package api

import (
	"net/http"
	"time"

	"github.com/trongbq/gotodo-server/internal"
	"github.com/trongbq/gotodo-server/internal/api/auth"
	"github.com/trongbq/gotodo-server/internal/database"
)

type userRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("UserID").(int64)
	user, err := s.db.GetUser(r.Context(), userID)
	if err == database.ErrNoRecordFound {
		s.respond(w, r, http.StatusNotFound, newErrResp(ErrCodeNotFound, "Current user not found"))
		return
	}
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	s.respond(w, r, http.StatusOK, user)
}

func (s *Server) registerUser(w http.ResponseWriter, r *http.Request) {
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
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	user.ID = id
	s.respond(w, r, http.StatusOK, user.User)
}

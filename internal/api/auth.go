package api

import (
	"net/http"

	"github.com/trongbq/gotodo-server/internal/api/auth"
	"github.com/trongbq/gotodo-server/internal/database"
)

type userSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authTokenResponse struct {
	Token string `json:"token"`
}

func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	var req userSignInRequest
	if err := s.decode(w, r, &req); err != nil {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, err.Error()))
		return
	}
	// Get user's id and password by username
	userID, userPasswd, err := s.db.GetUserIDAndPasswordByUsername(r.Context(), req.Username)
	if err != nil {
		if err == database.ErrNoRecordFound {
			s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, "Username or password is incorrect"))
			return
		}
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	// Compare existing password and request password
	if matched := auth.VerifyPassword(userPasswd, req.Password); !matched {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, "Username or password is incorrect"))
		return
	}
	// User valid, return response contains generated token
	token, err := s.auth.Generate(userID)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	s.respond(w, r, http.StatusOK, authTokenResponse{token})
}

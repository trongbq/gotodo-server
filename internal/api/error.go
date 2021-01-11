package api

import (
	"fmt"
	"time"
)

const (
	ErrCodeUnauthorized  = "UNAUTHORIZED"
	ErrCodeTokenExpired  = "TOKEN_EXPIRED"
	ErrCodeBadRequest    = "BAD_REQUEST"
	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeNotImplement  = "NOT_IMPLEMENT"
	ErrCodeInternalError = "INTERNAL_ERROR"
	ErrCodeUnknown       = "UNKNOWN_ERROR"
)

// Error struct for rest response
type ErrorResponse struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Info      map[string]interface{} `json:"info"`
	Timestamp time.Time              `json:"timestamp"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}

func newErrResp(code, msg string) ErrorResponse {
	return ErrorResponse{
		Code:      code,
		Message:   msg,
		Timestamp: time.Now(),
	}
}

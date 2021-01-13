package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/trongbq/gotodo-server/internal"
)

const (
	DefaultPage        = 1
	DefaultRowsPerPage = 10
)

type todoRequest struct {
	Content string `json:"content"`
}

type todoListResponse struct {
	Todos       []*internal.Todo `json:"todos"`
	Total       int64            `json:"total"`
	Page        int              `json:"page"`
	RowsPerPage int              `json:"rowsPerPage"`
}

func (s *Server) getTodoList(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("UserID").(int64)

	// Parse page and rowsPerPage query params from request
	var err error
	var page int
	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, "page query prams is invalid"))
			return
		}
	} else {
		page = DefaultPage
	}
	var rowsPerPage int
	rowsPerPageParam := r.URL.Query().Get("rowsPerPage")
	if rowsPerPageParam != "" {
		rowsPerPage, err = strconv.Atoi(rowsPerPageParam)
		if err != nil {
			s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, "rowsPerPage query prams is invalid"))
			return
		}
	} else {
		rowsPerPage = DefaultRowsPerPage
	}

	// Check value of total todos
	total, err := s.db.GetTodoCountByUser(r.Context(), userID)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	if total == 0 {
		s.respond(w, r, http.StatusOK, todoListResponse{
			Todos:       make([]*internal.Todo, 0),
			Total:       total,
			Page:        page,
			RowsPerPage: rowsPerPage,
		})
		return
	}
	// Get list of todos based on value of page and rowsPerPage
	todos, err := s.db.GetTodosByUser(r.Context(), userID, (page-1)*rowsPerPage, rowsPerPage)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	s.respond(w, r, http.StatusOK, todoListResponse{
		Todos:       todos,
		Total:       total,
		Page:        page,
		RowsPerPage: rowsPerPage,
	})
}

func (s *Server) addTodo(w http.ResponseWriter, r *http.Request) {
	var req todoRequest
	if err := s.decode(w, r, &req); err != nil {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, err.Error()))
		return
	}
	userID, _ := r.Context().Value("UserID").(int64)
	todo := internal.Todo{
		Content:   req.Content,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := s.db.InsertTodo(r.Context(), todo)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	todo.ID = id
	s.respond(w, r, http.StatusOK, todo)
}

func (s *Server) updateTodo(w http.ResponseWriter, r *http.Request) {
	var req todoRequest
	if err := s.decode(w, r, &req); err != nil {
		s.respond(w, r, http.StatusBadRequest, newErrResp(ErrCodeBadRequest, err.Error()))
		return
	}
	todoID, _ := strconv.ParseInt(chi.URLParam(r, "todoID"), 10, 64)
	err := s.db.UpdateTodoContent(r.Context(), todoID, req.Content)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
	}
	s.respond(w, r, http.StatusNoContent, nil)
}

func (s *Server) completeTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := strconv.ParseInt(chi.URLParam(r, "todoID"), 10, 64)
	err := s.db.UpdateTodoComplete(r.Context(), todoID)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	s.respond(w, r, http.StatusNoContent, nil)
}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := strconv.ParseInt(chi.URLParam(r, "todoID"), 10, 64)
	err := s.db.DeleteTodo(r.Context(), todoID)
	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, newErrResp(ErrCodeInternalError, err.Error()))
		return
	}
	s.respond(w, r, http.StatusNoContent, nil)
}

package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (s *Server) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Errorf("can not encode response data: %s", err)
		}
	}
}

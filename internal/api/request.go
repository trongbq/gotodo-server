package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	return d.Decode(v)
}

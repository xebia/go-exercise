package server

import (
	"net/http"
)

//nolint:revive
func (s *Server) createPatient(w http.ResponseWriter, r *http.Request) {
	writeError(w, "not implemented", http.StatusNotImplemented)
}

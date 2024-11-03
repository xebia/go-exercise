package server

import (
	"fmt"
	"net/http"
)

func (s *Server) createPatient(w http.ResponseWriter, r *http.Request) {
	writeError(w, fmt.Sprintf("not implemented"), http.StatusNotImplemented)
}

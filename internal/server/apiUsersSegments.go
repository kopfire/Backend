package server

import (
	"fmt"
	"net/http"
)

func (s *APIServer) handleUsersSegments(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUsers(w)
	case "POST":
		return s.handleCreateUser(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) linkSegmentsToUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package server

import (
	"AvitoBackend/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUsers(w)
	case "POST":
		return s.handleCreateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter) error {
	users, err := s.storage.GetUsers()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(domain.User)
	if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
		return err
	}

	user := domain.NewUserByName(createUserReq.Name)

	userResp, err := s.storage.CreateUser(user)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userResp)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}

	_, err = s.storage.GetUser(id)

	if err != nil {
		return fmt.Errorf("doesn't exists %s", strconv.Itoa(id))
	}

	if err := s.storage.DeleteUser(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": idStr})
}

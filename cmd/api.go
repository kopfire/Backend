package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/segment", makeHTTPHandleFunc(s.handleSegment))

	log.Println("Server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleSegment(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetSegment(w, r)
	case "POST":
		return s.handleCreateSegment(w, r)
	case "DELETE":
		return s.handleDeleteSegment(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)

}

func (s *APIServer) handleGetSegment(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateSegment(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteSegment(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) linkSegmentsToUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) getSegments(w http.ResponseWriter, r *http.Request) error {
	return nil
}

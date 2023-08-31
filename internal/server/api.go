package server

import (
	"AvitoBackend/internal/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	storage    storage.Storage
}

func NewAPIServer(listenAddr string, storage storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/segment", makeHTTPHandleFunc(s.handleSegment))
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/linkSegmentsToUser", makeHTTPHandleFunc(s.handleUsersSegments))
	router.HandleFunc("/getActiveSegmentsFromUser", makeHTTPHandleFunc(s.handleUsersSegments))

	log.Println("Server running on port: ", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		return
	}
}

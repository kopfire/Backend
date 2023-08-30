package server

import (
	"AvitoBackend/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) handleSegment(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetSegments(w)
	case "POST":
		return s.handleCreateSegment(w, r)
	case "DELETE":
		return s.handleDeleteSegment(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleGetSegments(w http.ResponseWriter) error {
	segments, err := s.storage.GetSegments()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, segments)
}

func (s *APIServer) handleCreateSegment(w http.ResponseWriter, r *http.Request) error {
	createSegmentReq := new(domain.Segment)
	if err := json.NewDecoder(r.Body).Decode(&createSegmentReq); err != nil {
		return err
	}

	segment := domain.NewSegment(createSegmentReq.Name)

	segmentResp, err := s.storage.CreateSegment(segment)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, segmentResp)
}

func (s *APIServer) handleDeleteSegment(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")

	if err := s.storage.DeleteSegment(name); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": name})
}

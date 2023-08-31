package server

import (
	"AvitoBackend/internal/domain"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	createSegmentReq := new(domain.SegmentReqJSON)
	if err := json.NewDecoder(r.Body).Decode(&createSegmentReq); err != nil {
		return err
	}

	segmentGet, err := s.storage.GetSegment(createSegmentReq.Name)

	if err == nil {
		return fmt.Errorf("already exists id:%s, name:%s", strconv.Itoa(segmentGet.ID), segmentGet.Name)
	}

	segment := domain.NewSegment(createSegmentReq.Name)

	segmentResp, err := s.storage.CreateSegment(segment)

	if err != nil {
		return err
	}

	if createSegmentReq.UserPercent != 0 {
		users, err := s.storage.GetUsers()
		if err != nil {
			return err
		}
		countUser := len(users) * createSegmentReq.UserPercent / 100

		set := make(map[int]bool)
		for len(set) != countUser {
			set[rand.Intn(len(users))] = true
		}

		for k := range set {
			if err != nil {
				return err
			}
			refUsersSegments := domain.NewRefUsersSegments(users[k].ID, segmentResp.ID)

			err = s.storage.CreateRefSegmentToUser(refUsersSegments)
			if err != nil {
				return err
			}
		}

	}

	return WriteJSON(w, http.StatusOK, segmentResp)
}

func (s *APIServer) handleDeleteSegment(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")

	_, err := s.storage.GetSegment(name)

	if err != nil {
		return fmt.Errorf("doesn't exists %s", name)
	}

	if err := s.storage.DeleteSegment(name); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": name})
}

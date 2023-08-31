package server

import (
	"AvitoBackend/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *APIServer) handleUsersSegments(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetActiveSegmentsFromUser(w, r)
	case "POST":
		return s.handleLinkSegmentsToUser(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleGetActiveSegmentsFromUser(w http.ResponseWriter, r *http.Request) error {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}

	_, err = s.storage.GetUser(id)

	if err != nil {
		return fmt.Errorf("doesn't exists %s", strconv.Itoa(id))
	}

	refUsersSegment, err := s.storage.GetActiveSegmentByUserId(id)
	if err != nil {
		return err
	}

	if refUsersSegment == nil {
		return WriteJSON(w, http.StatusOK, map[string]string{"message": "User doesn't have segment"})
	}

	var refSegments []string

	for _, v := range refUsersSegment {
		segment, err := s.storage.GetSegmentById(v.SegmentId)
		if err != nil {
			return err
		}
		refSegments = append(refSegments, segment.Name)
	}

	return WriteJSON(w, http.StatusOK, refSegments)
}

func (s *APIServer) handleLinkSegmentsToUser(w http.ResponseWriter, r *http.Request) error {
	linkSegmentsToUserReq := new(LinkSegmentsToUserJSON)
	if err := json.NewDecoder(r.Body).Decode(&linkSegmentsToUserReq); err != nil {
		return err
	}

	_, err := s.storage.GetUser(linkSegmentsToUserReq.ID)

	if err != nil {
		user := domain.NewUser(linkSegmentsToUserReq.ID)
		err := s.storage.CreateUserById(user)

		if err != nil {
			return err
		}
	}

	linkSegmentsToUserResp := new(LinkSegmentsToUserJSON)

	linkSegmentsToUserResp.ID = linkSegmentsToUserReq.ID

	for _, v := range linkSegmentsToUserReq.SegmentsAdd {
		segment, err := s.storage.GetSegment(v)
		if err != nil {
			return err
		}
		refUsersSegments := domain.NewRefUsersSegments(linkSegmentsToUserReq.ID, segment.ID)

		getUsersSegments, err := s.storage.GetIdRefSegmentToUser(refUsersSegments)
		if err == nil && getUsersSegments.Status != false {
			linkSegmentsToUserResp.SegmentsAdd = append(linkSegmentsToUserResp.SegmentsAdd, "already exists active link segment "+v)
		} else {
			if err == nil && getUsersSegments.Status == false {
				err = s.storage.UpdateRefSegmentToUser(getUsersSegments)
				if err != nil {
					return err
				}
			} else {
				err = s.storage.CreateRefSegmentToUser(refUsersSegments)
				if err != nil {
					return err
				}
			}
			linkSegmentsToUserResp.SegmentsAdd = append(linkSegmentsToUserResp.SegmentsAdd, "success link segment "+v)
		}
	}

	for _, v := range linkSegmentsToUserReq.SegmentsDelete {
		segment, err := s.storage.GetSegment(v)
		if err != nil {
			return err
		}
		refUsersSegments := domain.NewRefUsersSegments(linkSegmentsToUserReq.ID, segment.ID)

		getUsersSegments, err := s.storage.GetIdRefSegmentToUser(refUsersSegments)
		if err != nil {
			linkSegmentsToUserResp.SegmentsDelete = append(linkSegmentsToUserResp.SegmentsDelete, "doesn't exists active link segment "+v)
		} else {
			err = s.storage.DeleteRefSegmentToUser(getUsersSegments)
			if err != nil {
				return err
			}
			linkSegmentsToUserResp.SegmentsDelete = append(linkSegmentsToUserResp.SegmentsDelete, "success unlink segment "+v)
		}
	}

	return WriteJSON(w, http.StatusOK, linkSegmentsToUserResp) //CHANGE RESPONSE
}

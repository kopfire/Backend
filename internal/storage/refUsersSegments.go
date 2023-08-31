package storage

import (
	"AvitoBackend/internal/domain"
	"fmt"
	"strconv"
	"time"
)

func (s *PostgresStorage) CreateRefUsersSegmentsTable() error {
	query := `CREATE TABLE if not exists users_segments (
    	id SERIAL, 
    	user_id INTEGER, 
    	segment_id INTEGER, 
    	status BOOLEAN,
    	date_add TIMESTAMP,
    	date_del TIMESTAMP,
    	CONSTRAINT fk_users
			FOREIGN KEY (user_id) 
			REFERENCES users (id) ON  DELETE  CASCADE,
    	CONSTRAINT fk_segment
			FOREIGN KEY (segment_id) 
			REFERENCES segment (id) ON  DELETE  CASCADE
    )`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateRefSegmentToUser(ref *domain.RefUsersSegments) error {
	query := `INSERT INTO users_segments (user_id, segment_id, status, date_add) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Query(query, ref.UserId, ref.SegmentId, true, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) DeleteRefSegmentToUser(ref *domain.RefUsersSegments) error {
	query := `UPDATE users_segments SET status = $1, date_del = $2 WHERE id = $3`
	_, err := s.db.Query(query, false, time.Now().UTC(), ref.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) UpdateRefSegmentToUser(ref *domain.RefUsersSegments) error {
	query := `UPDATE users_segments SET status = $1, date_add = $2, date_del = $3 WHERE id = $4`
	_, err := s.db.Query(query, true, time.Now().UTC(), nil, ref.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) GetActiveSegmentByUserId(id int) ([]*domain.RefUsersSegments, error) {
	query := `SELECT * FROM users_segments WHERE user_id = $1 AND status = $2`
	rows, err := s.db.Query(query, id, true)
	if err != nil {
		return nil, err
	}
	var refUsersSegments []*domain.RefUsersSegments
	for rows.Next() {
		refUsersSegment := new(domain.RefUsersSegments)
		if err := rows.Scan(&refUsersSegment.Id,
			&refUsersSegment.UserId,
			&refUsersSegment.SegmentId,
			&refUsersSegment.Status,
			&refUsersSegment.DateAdd,
			&refUsersSegment.DateDel); err != nil {
			return nil, err
		}
		refUsersSegments = append(refUsersSegments, refUsersSegment)
	}
	return refUsersSegments, nil
}

func (s *PostgresStorage) GetIdRefSegmentToUser(ref *domain.RefUsersSegments) (*domain.RefUsersSegments, error) {
	rows, err := s.db.Query("SELECT * FROM users_segments WHERE user_id = $1 AND segment_id = $2", ref.UserId, ref.SegmentId)
	if err != nil {
		return nil, err
	}

	refUsersSegments := new(domain.RefUsersSegments)
	for rows.Next() {
		if err := rows.Scan(&refUsersSegments.Id,
			&refUsersSegments.UserId,
			&refUsersSegments.SegmentId,
			&refUsersSegments.Status,
			&refUsersSegments.DateAdd,
			&refUsersSegments.DateDel); err != nil {
			return nil, err
		}
	}

	if refUsersSegments.SegmentId != ref.SegmentId {
		return nil, fmt.Errorf("doesn't exists %s and %s",
			strconv.Itoa(ref.SegmentId),
			strconv.Itoa(ref.UserId))
	}

	return refUsersSegments, nil
}

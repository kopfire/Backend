package storage

import (
	"AvitoBackend/internal/domain"
	"fmt"
	"strconv"
)

func (s *PostgresStorage) CreateSegmentTable() error {
	query := `CREATE TABLE if not exists segment (
    	id SERIAL PRIMARY KEY,
        name VARCHAR(50))`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateSegment(segment *domain.Segment) (*domain.Segment, error) {
	var lastInsertId int
	query := `INSERT INTO segment (name) VALUES ($1) RETURNING id`
	err := s.db.QueryRow(query, segment.Name).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	segment.ID = lastInsertId

	return segment, err
}

func (s *PostgresStorage) DeleteSegment(name string) error {
	_, err := s.db.Query("DELETE FROM segment WHERE name = $1", name)
	return err
}

func (s *PostgresStorage) GetSegment(name string) (*domain.Segment, error) {
	rows, err := s.db.Query("SELECT * FROM segment WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	segment := new(domain.Segment)
	for rows.Next() {
		if err := rows.Scan(&segment.ID, &segment.Name); err != nil {
			return nil, err
		}
	}

	if segment.Name != name {
		return nil, fmt.Errorf("doesn't exists %s", name)
	}

	return segment, nil
}

func (s *PostgresStorage) GetSegmentById(id int) (*domain.Segment, error) {
	rows, err := s.db.Query("SELECT * FROM segment WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	segment := new(domain.Segment)
	for rows.Next() {
		if err := rows.Scan(&segment.ID, &segment.Name); err != nil {
			return nil, err
		}
	}

	if segment.ID != id {
		return nil, fmt.Errorf("doesn't exists %s", strconv.Itoa(id))
	}

	return segment, nil
}

func (s *PostgresStorage) GetSegments() ([]*domain.Segment, error) {
	rows, err := s.db.Query("SELECT * FROM segment")
	if err != nil {
		return nil, err
	}

	var segments []*domain.Segment
	for rows.Next() {
		segment := new(domain.Segment)
		if err := rows.Scan(&segment.ID, &segment.Name); err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

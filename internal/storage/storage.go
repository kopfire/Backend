package storage

import (
	"AvitoBackend/internal/domain"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Storage interface {
	CreateSegment(*domain.Segment) (*domain.Segment, error)
	DeleteSegment(string) error
	GetSegments() ([]*domain.Segment, error)

	CreateUser(*domain.User) (*domain.User, error)
	DeleteUser(int) error
	GetUsers() ([]*domain.User, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	err := s.createSegmentTable()
	if err != nil {
		return err
	}
	err = s.createUserTable()
	if err != nil {
		return err
	}
	return nil
}

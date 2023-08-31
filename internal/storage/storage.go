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
	GetSegment(string) (*domain.Segment, error)
	GetSegmentById(int) (*domain.Segment, error)

	CreateUser(*domain.User) (*domain.User, error)
	CreateUserById(*domain.User) error
	DeleteUser(int) error
	GetUsers() ([]*domain.User, error)
	GetUser(int) (*domain.User, error)

	CreateRefSegmentToUser(*domain.RefUsersSegments) error
	UpdateRefSegmentToUser(*domain.RefUsersSegments) error
	DeleteRefSegmentToUser(*domain.RefUsersSegments) error
	GetIdRefSegmentToUser(*domain.RefUsersSegments) (*domain.RefUsersSegments, error)
	GetActiveSegmentByUserId(int) ([]*domain.RefUsersSegments, error)
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
	err := s.CreateSegmentTable()
	if err != nil {
		return err
	}
	err = s.CreateUserTable()
	if err != nil {
		return err
	}
	err = s.CreateRefUsersSegmentsTable()
	if err != nil {
		return err
	}
	return nil
}

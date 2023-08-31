package storage

import (
	"AvitoBackend/internal/domain"
	"fmt"
	"strconv"
)

func (s *PostgresStorage) CreateUserTable() error {
	query := `CREATE TABLE if not exists users (
    	id SERIAL PRIMARY KEY,
        name VARCHAR(50))`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateUser(user *domain.User) (*domain.User, error) {
	var lastInsertId int
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err := s.db.QueryRow(query, user.Name).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	user.ID = lastInsertId

	return user, err
}

func (s *PostgresStorage) CreateUserById(user *domain.User) error {
	query := `INSERT INTO users (id, name) VALUES ($1, $2)`
	_, err := s.db.Query(query, user.ID, user.Name)
	return err
}

func (s *PostgresStorage) DeleteUser(id int) error {
	_, err := s.db.Query("DELETE FROM users WHERE id = $1", id)
	return err
}

func (s *PostgresStorage) GetUser(id int) (*domain.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	user := new(domain.User)
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
	}

	if user.ID != id {
		return nil, fmt.Errorf("doesn't exists %s", strconv.Itoa(id))
	}

	return user, nil
}

func (s *PostgresStorage) GetUsers() ([]*domain.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []*domain.User
	for rows.Next() {
		user := new(domain.User)
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

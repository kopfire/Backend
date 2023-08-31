package domain

import "strconv"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(id int) *User {
	return &User{
		ID:   id,
		Name: "test" + strconv.Itoa(id),
	}
}

func NewUserByName(name string) *User {
	return &User{
		Name: name,
	}
}

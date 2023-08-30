package domain

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUserEmpty() *User {
	return &User{}
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

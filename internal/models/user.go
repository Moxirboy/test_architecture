package models

type User struct {
	ID       string
	Password string
	Email    string
	Role
	At
}

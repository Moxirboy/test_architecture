package models

import "time"

type At struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
type Role string

const (
	RoleUser   Role = "User"
	RoleDoctor Role = "Doctor"
)
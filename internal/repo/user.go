package repo

import (
	"arch/internal/models"
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (err error)
	ExistUser(ctx context.Context, email string) (exist bool, err error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateIsVerified(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

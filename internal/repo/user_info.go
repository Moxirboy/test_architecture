package repo

import (
	"arch/internal/models"
	"context"
)
type IUserInfoRepository interface {
	CreateUserInfo(ctx context.Context, info *models.UserInfo) error
}
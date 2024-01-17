package postgres

import (
	"arch/internal/models"
	"arch/internal/repo"
	"arch/pkg/logger"
	"context"
	"database/sql"

)

type userInfoRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewUserInfoRepository(db *sql.DB, log logger.Logger) repo.IUserInfoRepository {
	return &userInfoRepository{
		db:  db,
		log: log,
	}
}
func (r *userInfoRepository) CreateUserInfo(ctx context.Context, info *models.UserInfo) error {
	err := r.db.QueryRowContext(
		ctx,
		createUserInfo,
		info.UserId,
		info.Name,
		info.Weigh,
		info.Height,
		info.Age,
		info.Waist,
		info.CreatedAt,
		info.Gender,
	).Scan(&info.Id)
	if err != nil {
		r.log.Error("repo.create.user error:",err)
		return err
	}
	return nil
}

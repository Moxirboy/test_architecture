package postgres

import (
	"arch/internal/models"
	"arch/internal/repo"
	"context"
	"database/sql"
	"log"
)

type userInfoRepository struct {
	db *sql.DB
	log log.Logger
}
func NewUserInfoRepository(db *sql.DB,log log.Logger) repo.IUserInfoRepository{
	return &userInfoRepository{
		db: db,
		log: log,
	}
}
func (r *userInfoRepository) CreateUserInfo(ctx context.Context,info models.UserInfo) error{
	err:=r.db.QueryRowContext(
		ctx,
		createUserInfo,
		
		)
}
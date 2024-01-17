package postgres

import (
	"arch/internal/models"
	"arch/internal/repo"
	"context"
	"database/sql"
	"log"
)

type userRepository struct {
	db  *sql.DB
	log log.Logger
}

func NewUserRepository(
	db *sql.DB,
	log log.Logger,
) repo.IUserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (err error) {
	err = r.db.QueryRowContext(
		ctx,
		createUser,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		false,
	).Scan(&user.ID)
	if err != nil {
		//r.log.Error("repo.User.create error:", err)
		return err
	}
	return nil
}
func (r *userRepository) ExistUser(ctx context.Context, email string) (exist bool, err error) {
	err = r.db.QueryRowContext(
		ctx,
		getUser,
		email,
	).Scan(&exist)
	if err != nil {
		return exist, err
	}
	return exist, err
}
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRowContext(
		ctx,
		getUserByEmail,
		email,
	).Scan(
		&user.ID,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateIsVerified(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(
		ctx,
		id,
	)
	if err != nil {
		// r.log.Error("repo.user.update error :",err)
		return err
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err != nil {
		//	r.log.Error("repo.user.delete error while transaction begin:", err)
		return err
	}
	_, err = tx.ExecContext(
		ctx,
		deleteUserById,
		id,
	)
	if err != nil {
		//	r.log.Error("repo.user.delete error :" ,err)
		_ = tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		//	r.log.Error("repo.box.create error while tx commit:", err.Error())
		return err
	}
	return nil
}

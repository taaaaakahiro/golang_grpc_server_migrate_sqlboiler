package persistence

import (
	"context"
	"database/sql"
	errs "github.com/friendsofgo/errors"
	"golang_grpc_proto/pkg/domain/models"
	"golang_grpc_proto/pkg/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserPersistence(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Get(ctx context.Context, id int) (*models.User, error) {
	query := `
	SELECT
	   id,
	   name
	FROM
	   users
	WHERE
		id = $1`

	stmtOut, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errs.WithStack(err)
	}
	var user models.User
	err = stmtOut.QueryRowContext(ctx, id).Scan(&user.ID, &user.Name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			break
		default:
			return nil, errs.WithStack(err)
		}
	}

	return &user, nil
}

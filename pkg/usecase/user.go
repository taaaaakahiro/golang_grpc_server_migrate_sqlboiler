package usecase

import (
	"context"
	errs "github.com/friendsofgo/errors"
	"golang_grpc_proto/pkg/domain/models"
	"golang_grpc_proto/pkg/infra/persistence"
)

type UserUseCase struct {
	repo *persistence.Repository
}

func NewUserUseCase(r *persistence.Repository) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) Get(ctx context.Context, id int) (*models.User, error) {
	user, err := uc.repo.User.Get(ctx, id)
	if err != nil {
		return nil, errs.WithStack(err)
	}
	if user == nil {
		return nil, nil
	}

	return user, nil
}

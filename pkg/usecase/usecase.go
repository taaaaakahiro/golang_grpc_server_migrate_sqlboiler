package usecase

import (
	"golang_grpc_proto/pkg/infra/persistence"
)

type UseCase struct {
	User *UserUseCase
}

func NewUseCase(r *persistence.Repository) *UseCase {
	return &UseCase{
		User: NewUserUseCase(r),
	}
}

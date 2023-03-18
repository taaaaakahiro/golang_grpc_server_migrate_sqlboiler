package repository

import (
	"context"
	"golang_grpc_proto/pkg/domain/models"
)

type IUserRepository interface {
	Get(ctx context.Context, id int) (*models.User, error)
}

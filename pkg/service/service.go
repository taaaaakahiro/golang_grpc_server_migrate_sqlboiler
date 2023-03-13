package service

import (
	"golang.org/x/exp/slog"
	"golang_grpc_proto/pkg/usecase"
)

type Service struct {
	logger      *slog.Logger
	UserService *UserService
}

func NewService(logger *slog.Logger, uc *usecase.UseCase) *Service {
	return &Service{
		UserService: NewUserService(logger, uc),
	}
}

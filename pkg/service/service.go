package service

import (
	"golang.org/x/exp/slog"
	"grpc_func_from_prcivate_repo/pkg/usecase"
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

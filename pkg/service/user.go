package service

import (
	"context"
	pb "github.com/taaaaakahiro/golang_grpc_proto/pb/proto"
	"golang.org/x/exp/slog"
	"grpc_func_from_prcivate_repo/pkg/usecase"
)

type UserService struct {
	logger *slog.Logger
	uc     usecase.UseCase
	pb.UnimplementedUserServiceServer
}

func NewUserService(logger *slog.Logger, uc *usecase.UseCase) *UserService {
	return &UserService{
		logger: logger,
		uc:     *uc,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.RegisterClientRequest) (*pb.RegisterClientResponse, error) {
	return &pb.RegisterClientResponse{
		ClientId: "create success!!!",
	}, nil

}

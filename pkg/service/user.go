package service

import (
	"context"
	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
	"golang.org/x/exp/slog"
	"golang_grpc_proto/pkg/usecase"
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

func (s *UserService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Id:   req.GetId(),
		Name: "get success",
	}, nil

}

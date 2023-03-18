package service

import (
	"context"
	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
	"golang.org/x/exp/slog"
	"golang_grpc_proto/pkg/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
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
	id, err := strconv.Atoi(req.GetId())
	if err != nil {
		s.logger.Error("failed to convert id", err)
		return nil, status.Error(codes.Internal, "unknown error occurred")
	}

	user, err := s.uc.User.Get(ctx, id)
	if err != nil {
		s.logger.Error("failed to get user", err)
		return nil, status.Error(codes.Internal, "unknown error occurred")
	}
	if user.ID == 0 {
		s.logger.Error("user is not found", err)
		return nil, status.Error(codes.NotFound, "user is not found")
	}

	return &pb.GetResponse{
		Id:   strconv.Itoa(user.ID),
		Name: user.Name,
	}, nil

}

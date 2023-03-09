package service

import (
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

type Service struct {
	GrpcServ *grpc.Server
}

func NewService(logger *slog.Logger) *Service {
	s := grpc.NewServer()

	return &Service{
		GrpcServ: s,
	}
}

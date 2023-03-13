package server

import (
	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
	"golang_grpc_proto/pkg/service"
	"google.golang.org/grpc"
)

type Server struct {
	GrpcServ *grpc.Server
}

func NewServer(services *service.Service) *Server {
	s := &Server{
		GrpcServ: grpc.NewServer(),
	}
	s.registerService(services)

	return s
}

func (s *Server) registerService(services *service.Service) {
	pb.RegisterUserServiceServer(s.GrpcServ, services.UserService)

}

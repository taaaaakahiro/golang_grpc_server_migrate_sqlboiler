package server

import (
	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
	"golang_grpc_proto/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
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
	s.healthCheck()
	pb.RegisterUserServiceServer(s.GrpcServ, services.UserService)

}

func (s *Server) healthCheck() {
	reflection.Register(s.GrpcServ)
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(s.GrpcServ, healthSrv)
	healthSrv.SetServingStatus("golang_grpc_server_api", healthpb.HealthCheckResponse_SERVING)
}

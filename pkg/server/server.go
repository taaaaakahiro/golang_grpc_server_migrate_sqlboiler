package server

import (
	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
	"golang_grpc_proto/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const healthCheckStatus = "golang_grpc_server_api"

type Server struct {
	GrpcServ *grpc.Server
}

func NewServer(services *service.Service) *Server {
	s := &Server{
		GrpcServ: grpc.NewServer(),
	}
	reflection.Register(s.GrpcServ)
	s.registerService(services)

	return s
}

func (s *Server) registerService(services *service.Service) {
	s.healthCheck()
	pb.RegisterUserServiceServer(s.GrpcServ, services.UserService)
}

func (s *Server) healthCheck() {
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(s.GrpcServ, healthSrv)
	healthSrv.SetServingStatus(healthCheckStatus, healthpb.HealthCheckResponse_SERVING)
}

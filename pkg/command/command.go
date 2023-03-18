package command

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"golang_grpc_proto/pkg/config"
	"golang_grpc_proto/pkg/infra/persistence"
	"golang_grpc_proto/pkg/io"
	"golang_grpc_proto/pkg/server"
	"golang_grpc_proto/pkg/service"
	"golang_grpc_proto/pkg/usecase"
	"net"
	"os"
	"os/signal"
)

const (
	exitOk = 0
	exitNG = 1
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	// init logger
	logger := slog.New(slog.NewTextHandler(os.Stderr))

	// init config
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", err)
		return exitNG
	}

	db, err := io.NewDataBase(cfg)
	if err != nil {
		logger.Error("failed to connect database", err)
		return exitNG
	}

	repositories := persistence.NewRepository(db)

	// init Listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Error("failed to listen port", err)
		return exitNG
	}

	// init UseCase
	useCases := usecase.NewUseCase(repositories)

	// init Service
	services := service.NewService(logger, useCases)

	// init server
	server := server.NewServer(services)

	// run grpc server
	go func() {
		logger.InfoCtx(ctx, "Starting gRPC Server", "PORT", cfg.Server.Port)
		server.GrpcServ.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping gRPC server...")
	server.GrpcServ.GracefulStop()

	return exitOk
}

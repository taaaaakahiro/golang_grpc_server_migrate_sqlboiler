package command

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/reflection"
	"grpc_func_from_prcivate_repo/pkg/config"
	"grpc_func_from_prcivate_repo/pkg/service"
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

	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", err)
		return exitNG
	}

	// init lisner
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Error("failed to listen port", err)
		return exitNG
	}

	// init server
	s := service.NewService(logger)

	// reflectionを設定(grpcurlコマンドを受け付ける設定)
	reflection.Register(s.GrpcServ)

	// run grpc server
	go func() {
		logger.InfoCtx(ctx, "starting gRPC server...", "PORT", cfg.Server.Port)
		s.GrpcServ.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping gRPC server...")
	s.GrpcServ.GracefulStop()

	return exitOk
}

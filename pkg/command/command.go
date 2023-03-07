package command

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"grpc_func_from_prcivate_repo/pkg/config"
	"grpc_func_from_prcivate_repo/pkg/server"
	"net/http"
	"os"
	"os/signal"
	"time"

	pb "github.com/taaaaakahiro/golang_grpc_proto/pb/proto"
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

	// init grpc
	_ = pb.NewUserServiceClient(nil) //todo

	// init server
	s := server.NewServer()

	go func() {
		if err = s.Echo.Start(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			s.Echo.Logger.Fatal("shutting down the server")
		}
	}()

	// シグナルを待機してGraceful Shutdownを実行
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	s.Echo.Logger.Info("shutdown server...")

	// コンテキストのタイムアウトを設定してサーバーをシャットダウン
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err = s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
	}

	s.Echo.Logger.Info("server exiting")
	return exitOk
}

package server

import (
	"context"
	"golang.org/x/exp/slog"
	"golang_grpc_proto/pkg/config"
	"golang_grpc_proto/pkg/infra/persistence"
	"golang_grpc_proto/pkg/io"
	"golang_grpc_proto/pkg/service"
	"golang_grpc_proto/pkg/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"

	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
)

const bufSize = 1024 * 1024

var (
	lis  *bufconn.Listener
	conn *grpc.ClientConn
)

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stderr))
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db, err := io.NewDataBase(cfg)
	if err != nil {
		log.Fatal(err)
	}
	r := persistence.NewRepository(db)

	uc := usecase.NewUseCase(r)

	s := service.NewService(logger, uc)

	lis = bufconn.Listen(bufSize)
	// init server
	testServ := grpc.NewServer()
	pb.RegisterUserServiceServer(testServ, s.UserService)
	// run server
	go func() {
		if err := testServ.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	// conn server
	conn, err = grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	res := m.Run()
	// after

	os.Exit(res)

}

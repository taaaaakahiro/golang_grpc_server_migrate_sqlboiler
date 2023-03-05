package main

import (
	"context"
	"fmt"

	pb "github.com/taaaaakahiro/golang_grpc_proto/pb/proto"
	"github.com/taaaaakahiro/golang_private_repo/pkg"
)

func main() {
	fmt.Println("hello world")
	userServ := pb.NewUserServiceClient(nil)
	res, err := userServ.GetCurrentUserAccount(context.Background(), nil)
	if err != nil {
		return
	}
	pkg.GetPrivateRepo()
	res.GetClientId()
}

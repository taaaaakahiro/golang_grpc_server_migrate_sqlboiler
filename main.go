package main

import (
	"context"
	"fmt"

	userPb "github.com/taaaaakahiro/golang_grpc_proto/pb/proto"
)

func main() {
	fmt.Println("hello world")
	userServ := userPb.NewUserServiceClient(nil)
	res, err := userServ.GetCurrentUserAccount(context.Background(), nil)
	if err != nil {
		return
	}
	fmt.Println(res.GetClientId()) // とりあえず

}

package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	pb "github.com/taaaaakahiro/golang_grpc_proto/pkg/grpc"
	"testing"
)

func TestServer_GetUser(t *testing.T) {
	type args struct {
		c      context.Context
		userID string
	}

	tests := []struct {
		name   string
		args   args
		checks func(t *testing.T, got *pb.GetResponse, err error)
	}{
		{
			name: "ok",
			args: args{
				c:      context.Background(),
				userID: "1",
			},
			checks: func(t *testing.T, got *pb.GetResponse, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)

				assert.Equal(t, "1", got.Id)
				assert.Equal(t, "user1", got.Name)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pb.NewUserServiceClient(conn).Get(tt.args.c, &pb.GetRequest{Id: tt.args.userID})
			tt.checks(t, got, err)
		})
	}

}

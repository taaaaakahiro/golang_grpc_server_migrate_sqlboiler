package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"testing"
)

func TestServer_HealthCheck(t *testing.T) {
	type args struct {
		c   context.Context
		req *healthpb.HealthCheckRequest
	}

	tests := []struct {
		name   string
		args   args
		checks func(t *testing.T, got *healthpb.HealthCheckResponse, err error)
	}{
		{
			name: "ok",
			args: args{
				c:   context.Background(),
				req: &healthpb.HealthCheckRequest{Service: healthCheckStatus},
			},
			checks: func(t *testing.T, got *healthpb.HealthCheckResponse, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
				assert.Equal(t, "status:SERVING", got.String())

			},
		},
		{
			name: "ng: unknown service",
			args: args{
				c:   context.Background(),
				req: &healthpb.HealthCheckRequest{Service: "unknown"},
			},
			checks: func(t *testing.T, got *healthpb.HealthCheckResponse, err error) {
				assert.Error(t, err)
				assert.Empty(t, got)
				assert.Equal(t, "rpc error: code = NotFound desc = unknown service", err.Error())

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := healthSrv.Check(tt.args.c, tt.args.req)
			tt.checks(t, got, err)

		})
	}
}

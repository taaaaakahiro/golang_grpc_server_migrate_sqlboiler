package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server *serverConfig
}

type serverConfig struct {
	Port int `env:"PORT,required"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

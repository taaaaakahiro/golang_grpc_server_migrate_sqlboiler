package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server *serverConfig
	DB     *databaseConfig
}

type databaseConfig struct {
	User     string `env:"POSTGRES_USER,required"`
	Database string `env:"POSTGRES_DATABASE,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Host     string `env:"POSTGRES_HOST,required"`
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

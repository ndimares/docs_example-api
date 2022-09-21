package config

import (
	"context"

	"github.com/speakeasy-api/speakeasy-core/config"
	"github.com/speakeasy-api/speakeasy-core/drivers/psql"
	"github.com/speakeasy-api/speakeasy-core/listeners/http"
)

type SDK struct {
	APIKey string `env:"SPEAKEASY_SDK_API_KEY"`
}

type Config struct {
	Database psql.Config
	HTTP     http.Config `yaml:"http"`
	SDK      SDK
}

// Load loads the configuration from the config/config.yaml file.
func Load(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	if err := config.Load(ctx, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

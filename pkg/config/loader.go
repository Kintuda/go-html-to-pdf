package config

import (
	"os"

	"github.com/caarlos0/env"
	gotdotenv "github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(LoadConfigFromEnv[AppConfig]),
)

type Configs interface {
	AppConfig
}

func LoadConfigFromEnv[k Configs]() (*k, error) {
	var config k

	if os.Getenv("ENV") != "production" {
		if err := gotdotenv.Load(); err != nil {
			return &config, err
		}
	}

	if err := env.Parse(&config); err != nil {
		return &config, err
	}

	return &config, nil
}

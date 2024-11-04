package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Server         Server
	PostgresConfig Postgres
}

func Setup() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return config, errors.Wrap(err, "failed to process config")
	}

	return config, nil
}

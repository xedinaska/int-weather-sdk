package config

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

type Config struct{}

func Load() (*Config, error) {
	cfg := &Config{}
	for _, c := range []interface{}{
		cfg,
	} {
		if err := env.Parse(c); err != nil {
			return nil, errors.Wrapf(err, "env.Pars(..) failed to load configuration.")
		}
	}
	return cfg, nil
}

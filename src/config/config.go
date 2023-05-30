package config

import "github.com/caarlos0/env/v6"

type Config struct {
	PgURL string `env:"PgURL" envDefault:"user=postgres password=password dbname=postgres sslmode=disable host=db port=5432"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

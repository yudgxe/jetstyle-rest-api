package app

import "github.com/yudgxe/jetstyle-rest-api/pkg/database"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	DB       *database.PostgresConnInfo
}

func NewConfig() *Config {
	return &Config{}
}
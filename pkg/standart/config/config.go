package config

import (
	"github.com/wilian746/gorm-crud-generator/pkg/standart/utils/env"
)

type Config struct {
	Port int
	Timeout int
	Dialect string
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        env.GetEnvInt("PORT", 8666),
		Timeout:     env.GetEnvInt("TIMEOUT", 30),
		Dialect:     env.GetEnv("DATABASE_DIALECT", "sqlite3"),
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory:"),
	}
}

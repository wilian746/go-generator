package config

import "github.com/wilian746/gorm-crud-generator/pkg/standart/utils/get_env"

type Config struct {
	Port int
	Timeout int
	Dialect string
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        getenv.GetEnvInt("PORT", 8666),
		Timeout:     getenv.GetEnvInt("TIMEOUT", 30),
		Dialect:     getenv.GetEnv("DATABASE_DIALECT", "sqlite3"),
		DatabaseURI: getenv.GetEnv("DATABASE_URI", ":memory:"),
	}
}

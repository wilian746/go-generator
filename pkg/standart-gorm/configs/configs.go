package config

import "github.com/wilian746/go-generator/pkg/standart-gorm/internal/utils/environment"

type Config struct {
	Port        int
	Timeout     int
	Dialect     string
	DatabaseURI string
	SwaggerHost string
}

func GetConfig() Config {
	return Config{
		Port:        environment.GetEnvAndParseToInt("PORT", 8080),
		Timeout:     environment.GetEnvAndParseToInt("TIMEOUT", 30),
		Dialect:     environment.GetEnvString("DATABASE_DIALECT", "sqlite3"),
		DatabaseURI: environment.GetEnvString("DATABASE_URI", ":memory:"),
		SwaggerHost: environment.GetEnvString("SWAGGER_HOST", "localhost:8080"),
	}
}

package getenv

import (
	"os"
	"strconv"
)

func GetEnv(envName string, defaultValue string) string {
	environment := os.Getenv(envName)
	if environment == "" {
		return defaultValue
	}
	return environment
}

func GetEnvInt(envName string, defaultValue int) int {
	environment := os.Getenv(envName)
	if environment == "" {
		return defaultValue
	}
	environmentNum, err := strconv.Atoi(environment)
	if err != nil {
		return defaultValue
	}
	return environmentNum
}

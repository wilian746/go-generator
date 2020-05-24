package environment

import (
	"os"
	"strconv"
)

func GetEnvString(envName, defaultValue string) string {
	environment := os.Getenv(envName)
	if environment == "" {
		return defaultValue
	}
	return environment
}

func GetEnvAndParseToInt(envName string, defaultValue int) int {
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

package environment

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Run("Should get default value of env string", func(t *testing.T) {
		environment := GetEnvString("ENVIRONMENT", "default_env")
		assert.Equal(t, environment, "default_env")
	})
	t.Run("Should not get default value of env string", func(t *testing.T) {
		_ = os.Setenv("ENVIRONMENT", "other_env")
		environment := GetEnvString("ENVIRONMENT", "default_env")
		assert.Equal(t, environment, "other_env")
	})
}

func TestGetEnvInt(t *testing.T) {
	t.Run("Should get default value of env string", func(t *testing.T) {
		environment := GetEnvAndParseToInt("ENVIRONMENT", 123)
		assert.Equal(t, environment, 123)
	})
	t.Run("Should not get default value of env string", func(t *testing.T) {
		_ = os.Setenv("ENVIRONMENT", "987")
		environment := GetEnvAndParseToInt("ENVIRONMENT", 123)
		assert.Equal(t, environment, 987)
	})
}

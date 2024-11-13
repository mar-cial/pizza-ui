package adapters

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort_EnvVarSet(t *testing.T) {
	expectedPort := "9090"

	t.Run("set UI_PORT env variable", func(t *testing.T) {
		// Set the UI_PORT environment variable
		os.Setenv("UI_PORT", expectedPort)
	})

	config := NewEnvConfig()
	port, err := config.GetPort()

	t.Run("port should be set", func(t *testing.T) {
		assert.Nil(t, err, "Expected no error when UI_PORT is set")
		assert.Equal(t, ":"+expectedPort, port, "Expected port to match the environment variable")
	})

	defer os.Unsetenv("UI_PORT") // Clean up after the test
}

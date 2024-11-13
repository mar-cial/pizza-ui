package adapters

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort_EnvVarSet(t *testing.T) {
	expectedPort := "9090"
	expectedProjectName := "testing_name"

	portEnvKey := "UI_PORT"
	projectNameEnvKey := "PROJECT_NAME"

	config := NewEnvConfig()

	t.Run("set UI_PORT env variable", func(t *testing.T) {
		os.Setenv(portEnvKey, expectedPort)

		port, err := config.GetPort()
		assert.Nil(t, err, "no port found")
		assert.Equal(t, port, fmt.Sprintf(":%s", expectedPort), "port should be the same")

		if err := os.Unsetenv(portEnvKey); err != nil {
			t.Error(err)
		}
	})

	t.Run("set PROJECT_NAME env variable", func(t *testing.T) {
		os.Setenv(projectNameEnvKey, expectedProjectName)

		name, err := config.GetProjectName()

		assert.Nil(t, err, "no project name found")
		assert.Equal(t, name, expectedProjectName, "project name should be the same")
	})

	defer os.Unsetenv(portEnvKey) // Clean up after the test
	defer os.Unsetenv(projectNameEnvKey)
}

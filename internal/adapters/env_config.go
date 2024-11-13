package adapters

import (
	"errors"
	"fmt"
	"os"

	"github.com/mar-cial/pizza-ui/internal/ports"
)

var (
	ErrBadPort = errors.New("port variable needs to be set")
	ErrBadName = errors.New("")
)

type envConfig struct {
}

func (e *envConfig) GetPort() (string, error) {
	port := os.Getenv("UI_PORT")
	if port == "" {
		return "", ErrBadPort
	}

	port = fmt.Sprintf(":%s", port)

	return port, nil
}

func (e *envConfig) GetProjectName() (string, error) {
	projectName := os.Getenv("PROJECT_NAME")
	if projectName == "" {
		return "", ErrBadName
	}

	return projectName, nil
}

func NewEnvConfig() ports.Config {
	return &envConfig{}
}

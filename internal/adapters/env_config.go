package adapters

import (
	"errors"
	"fmt"
	"os"

	"github.com/mar-cial/pizza-ui/internal/ports"
)

var (
	ErrBadPort = errors.New("port variable needs to be set")
)

type envConfig struct {
}

func (e *envConfig) GetPort() (string, error) {
	port := os.Getenv("UI_PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port), nil
}

func NewEnvConfig() ports.Config {
	return &envConfig{}
}

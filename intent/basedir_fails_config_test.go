package intent_test

import (
	"errors"
)

type BasedirFailsConfig struct{}

func (c BasedirFailsConfig) GetBaseDir() (string, error) {
	return "", errors.New("expected!")
}

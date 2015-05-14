package intent_test

import (
	"errors"
	"github.com/moosingin3space/libintent.go/intent"
)

type BasedirFailsConfig struct {
	intent.BaseUnixConfiguration
}

func (c BasedirFailsConfig) GetBaseDir() (string, error) {
	return "", errors.New("expected!")
}

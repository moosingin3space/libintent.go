package intent_test

import (
	"errors"
)

type MkdirFailsBaseConfig struct{}

func (c MkdirFailsBaseConfig) GetBaseDir() (string, error) {
	return "/path", nil
}

func (c MkdirFailsBaseConfig) DirExists(path string) bool {
	return false
}

type MkdirFailsFirstConfig struct {
	MkdirFailsBaseConfig
}

func (c MkdirFailsFirstConfig) Mkdir(path string) error {
	if path == "/path/.intent" {
		return errors.New("expected!")
	}
	return nil
}

type MkdirFailsSecondConfig struct {
	MkdirFailsBaseConfig
}

func (c MkdirFailsSecondConfig) Mkdir(path string) error {
	if path == "/path/.intent/handler" {
		return errors.New("expected!")
	}
	return nil
}

type MkdirFailsThirdConfig struct {
	MkdirFailsBaseConfig
}

func (c MkdirFailsThirdConfig) Mkdir(path string) error {
	if path == "/path/.intent/comm" {
		return errors.New("expected!")
	}
	return nil
}

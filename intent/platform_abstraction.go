package intent

import (
	"io"
)

type Platform interface {
	Init() error
	Destroy() error

	Protocol(name string) string
	Conversation(name string) string

	NewSocket(name string) (io.ReadWriter, error)
	CleanupSocket(name string) error
}

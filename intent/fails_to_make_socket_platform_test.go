package intent_test

import (
	"io"
)

type FailsToMakeSocketPlatform struct{}

func (p FailsToMakeSocketPlatform) Init() error {
	return nil
}
func (p FailsToMakeSocketPlatform) Destroy() error {
	return nil
}
func (p FailsToMakeSocketPlatform) Protocol(name string) string {
	return name
}
func (p FailsToMakeSocketPlatform) Conversation(name string) string {
	return name
}
func (p FailsToMakeSocketPlatform) NewSocket(name string) (io.ReadWriter, error) {
	return nil, io.ErrUnexpectedEOF
}
func (p FailsToMakeSocketPlatform) CleanupSocket(name string) error {
	return nil
}

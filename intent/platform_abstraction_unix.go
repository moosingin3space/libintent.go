// +build darwin dragonfly freebsd linux netbsd openbsd

package intent

import (
	"io"
	"net"
	"os"
	osuser "os/user"
	"path/filepath"
)

const INTENT_DIRECTORY = ".intent"
const HANDLER_DIRECTORY = "handler"
const COMM_DIRECTORY = "comm"

func createUnixSocket(path string) (conn net.Conn, err error) {
	addr := &net.UnixAddr{Name: path, Net: "unixgram"}
	conn, err = net.ListenUnixgram("unixgram", addr)
	return
}

type Configuration interface {
	GetBaseDir() (string, error)
}

type UnixPlatform struct {
	Config Configuration
}

type UserUnixConfiguration struct{}

func (p UnixPlatform) Init() (err error) {
	// Create directories
	baseDir, err := p.Config.GetBaseDir()
	if err != nil {
		return
	}
	intentRootDir := filepath.Join(baseDir, INTENT_DIRECTORY)
	if _, e := os.Stat(intentRootDir); os.IsNotExist(e) {
		err = os.Mkdir(intentRootDir, 0700)
		if err != nil {
			return
		}
	}
	handlerDir := filepath.Join(intentRootDir, HANDLER_DIRECTORY)
	if _, e := os.Stat(handlerDir); os.IsNotExist(e) {
		err = os.Mkdir(handlerDir, 0700)
		if err != nil {
			return
		}
	}
	commDir := filepath.Join(intentRootDir, COMM_DIRECTORY)
	if _, e := os.Stat(commDir); os.IsNotExist(e) {
		err = os.Mkdir(commDir, 0700)
		if err != nil {
			return
		}
	}
	return
}

func (p UnixPlatform) Destroy() (err error) {
	baseDir, err := p.Config.GetBaseDir()
	if err != nil {
		return
	}

	intentRootDir := filepath.Join(baseDir, INTENT_DIRECTORY)
	os.RemoveAll(intentRootDir)
	return
}

func (p UnixPlatform) Protocol(name string) string {
	return filepath.Join(HANDLER_DIRECTORY, name)
}

func (p UnixPlatform) Conversation(name string) string {
	return filepath.Join(COMM_DIRECTORY, name)
}

func (p UnixPlatform) NewSocket(name string) (conn io.ReadWriter, err error) {
	baseDir, err := p.Config.GetBaseDir()
	if err != nil {
		return
	}

	path := filepath.Join(baseDir, INTENT_DIRECTORY, name)

	conn, err = createUnixSocket(path)
	return
}

func (p UnixPlatform) CleanupSocket(name string) (err error) {
	baseDir, err := p.Config.GetBaseDir()
	if err != nil {
		return
	}

	path := filepath.Join(baseDir, INTENT_DIRECTORY, name)
	os.Remove(path)
	return
}

func (c UserUnixConfiguration) GetBaseDir() (string, error) {
	user, err := osuser.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}

func DefaultUnixPlatform() Platform {
	return UnixPlatform{Config: UserUnixConfiguration{}}
}

func DefaultPlatform() Platform {
	return DefaultUnixPlatform()
}

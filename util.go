package intent

// TODO think about another abstraction
// that will allow libintent.go to work
// on other OSes
import (
	"net"
	"os"
	osuser "os/user"
	"path/filepath"
)

const INTENT_DIRECTORY = ".intent"
const HANDLER_DIRECTORY = "handler"
const COMM_DIRECTORY = "comm"

func makeIntentDirectories() (err error) {
	user, err := osuser.Current()
	if err != nil {
		return
	}
	intentRootDir := filepath.Join(user.HomeDir, INTENT_DIRECTORY)
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

func makeProtocolSocket(protocol string) (conn net.Conn, err error) {
	user, err := osuser.Current()
	if err != nil {
		return
	}

	path := filepath.Join(user.HomeDir,
		INTENT_DIRECTORY, HANDLER_DIRECTORY, protocol)
	addr := &net.UnixAddr{path, "unixgram"}
	conn, err = net.ListenUnixgram("unixgram", addr)
	if err != nil {
		return
	}
	return
}

func removeProtocolSocket(protocol string) (err error) {
	user, err := osuser.Current()
	if err != nil {
		return
	}

	path := filepath.Join(user.HomeDir,
		INTENT_DIRECTORY, HANDLER_DIRECTORY, protocol)
	os.Remove(path)
	return
}

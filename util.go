package intent

// TODO think about another abstraction
// that will allow libintent.go to work
// on windows or plan9
import (
	"errors"
	sys "golang.org/x/sys/unix"
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
	err = os.Mkdir(intentRootDir, 0700)
	if err != nil {
		return
	}
	handlerDir := filepath.Join(intentRootDir, HANDLER_DIRECTORY)
	err = os.Mkdir(handlerDir, 0700)
	if err != nil {
		return
	}
	commDir := filepath.Join(intentRootDir, COMM_DIRECTORY)
	err = os.Mkdir(commDir, 0700)
	if err != nil {
		return
	}
	return
}

func makeProtocolSocket(protocol string) (path string, err error) {
	user, err := osuser.Current()
	if err != nil {
		return
	}
	path := filepath.Join(user.HomeDir,
		INTENT_DIRECTORY, HANDLER_DIRECTORY, protocol)
	err = sys.Mkfifo(path, 0600)
	if err != nil {
		return
	}
	return
}

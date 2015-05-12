// +build darwin dragonfly freebsd linux netbsd openbsd

package intent

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	osuser "os/user"
	"path/filepath"
	"testing"
)

func TestInitAndDestroy(t *testing.T) {
	Convey("Given a UNIX platform", t, func() {
		platform := UnixPlatform{}
		user, err := osuser.Current()
		So(err, ShouldBeNil)

		Convey("Init should create a set of directories", func() {
			err := platform.Init()
			So(err, ShouldBeNil)

			intentRootDir := filepath.Join(user.HomeDir, INTENT_DIRECTORY)
			_, err = os.Stat(intentRootDir)
			So(err, ShouldBeNil)
			So(os.IsNotExist(err), ShouldBeFalse)

			handlerDir := filepath.Join(intentRootDir, HANDLER_DIRECTORY)
			_, err = os.Stat(handlerDir)
			So(err, ShouldBeNil)
			So(os.IsNotExist(err), ShouldBeFalse)

			commDir := filepath.Join(intentRootDir, COMM_DIRECTORY)
			_, err = os.Stat(commDir)
			So(err, ShouldBeNil)
			So(os.IsNotExist(err), ShouldBeFalse)
		})

		Convey("Destroy should delete them", func() {
			err := platform.Destroy()
			So(err, ShouldBeNil)

			intentRootDir := filepath.Join(user.HomeDir, INTENT_DIRECTORY)
			_, err = os.Stat(intentRootDir)
			So(err, ShouldNotBeNil)
			So(os.IsNotExist(err), ShouldBeTrue)
		})
	})
}

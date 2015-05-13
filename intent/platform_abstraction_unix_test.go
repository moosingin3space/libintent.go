// +build darwin dragonfly freebsd linux netbsd openbsd

package intent_test

import (
	"github.com/moosingin3space/libintent.go/intent"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	osuser "os/user"
	"path/filepath"
	"testing"
)

func TestInitAndDestroy(t *testing.T) {
	Convey("Given a default UNIX platform", t, func() {
		platform := intent.DefaultUnixPlatform()
		user, err := osuser.Current()
		So(err, ShouldBeNil)

		Convey("Init should create a set of directories", func() {
			err := platform.Init()
			So(err, ShouldBeNil)

			intentRootDir := filepath.Join(user.HomeDir, ".intent")
			_, err = os.Stat(intentRootDir)
			So(err, ShouldBeNil)
			So(os.IsNotExist(err), ShouldBeFalse)

			handlerDir := filepath.Join(intentRootDir, "handler")
			_, err = os.Stat(handlerDir)
			So(err, ShouldBeNil)
			So(os.IsNotExist(err), ShouldBeFalse)

			commDir := filepath.Join(intentRootDir, "comm")
			_, err = os.Stat(commDir)
			So(err, ShouldBeNil)
			So(os.IsNotExist(err), ShouldBeFalse)
		})

		Convey("Destroy should delete them", func() {
			err := platform.Destroy()
			So(err, ShouldBeNil)

			intentRootDir := filepath.Join(user.HomeDir, ".intent")
			_, err = os.Stat(intentRootDir)
			So(err, ShouldNotBeNil)
			So(os.IsNotExist(err), ShouldBeTrue)
		})
	})
}

func TestHandlesErrors(t *testing.T) {
	Convey("Given a Unix platform without a basedir", t, func() {
		platform := intent.UnixPlatform{Config: BasedirFailsConfig{}}

		expectToFail := func(err error) {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "expected!")
		}

		Convey("Init should pass the expected error down", func() {
			expectToFail(platform.Init())
		})

		Convey("Destroy should pass the expected error down", func() {
			expectToFail(platform.Destroy())
		})

		Convey("Creating new sockets should error", func() {
			_, err := platform.NewSocket("sock")
			expectToFail(err)
		})

		Convey("Cleaning up sockets should error", func() {
			expectToFail(platform.CleanupSocket("sock"))
		})
	})
}

func TestPathnamesAreExpected(t *testing.T) {
	Convey("Given a UNIX platform", t, func() {
		platform := intent.DefaultUnixPlatform()
		Convey("Protocols should use the path handler/protocol", func() {
			protocol := platform.Protocol("http")
			So(protocol, ShouldEqual, "handler/http")
		})

		Convey("Conversations should use the path comm/protocol", func() {
			comm := platform.Conversation("conv1")
			So(comm, ShouldEqual, "comm/conv1")
		})
	})
}

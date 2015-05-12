package intent

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func WithPlatform(platform Platform, f func(p Platform)) func() {
	return func() {
		err := platform.Init()
		So(err, ShouldBeNil)

		if f != nil {
			f(platform)
		}

		err = platform.Destroy()
		So(err, ShouldBeNil)
	}
}

func TestIntentRegistration(t *testing.T) {
	t.SkipNow()
}

func TestDefaultPlatform(t *testing.T) {
	Convey("Given the default Platform", t, func() {
		app := Application{
			Name:    "Awesome Listing Tool",
			Version: "0.0.1",
		}
		handler := make(chan Intent)
		validator := func(intent Intent) bool {
			return true
		}

		Convey("Default platform initialization and destruction will not fail", WithPlatform(DefaultPlatform(), nil))

		Convey("Intent registration and unregistration will not fail", WithPlatform(DefaultPlatform(), func(platform Platform) {
			recv, err := Register(platform, "ls", app, validator, handler)

			So(err, ShouldBeNil)
			So(recv, ShouldNotBeNil)

			Unregister(platform, recv)
		}))
	})
}

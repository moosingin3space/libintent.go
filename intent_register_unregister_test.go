package intent_test

import (
	intent "github.com/moosingin3space/libintent.go"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func WithPlatform(platform intent.Platform, f func(p intent.Platform)) func() {
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
		app := intent.Application{
			Name:    "Awesome Listing Tool",
			Version: "0.0.1",
		}
		handler := make(chan intent.Intent)
		validator := func(intent intent.Intent) bool {
			return true
		}

		Convey("Default platform initialization and destruction will not fail", WithPlatform(intent.DefaultPlatform(), nil))

		Convey("Intent registration and unregistration will not fail", WithPlatform(intent.DefaultPlatform(), func(platform intent.Platform) {
			recv, err := intent.Register(platform, "ls", app, validator, handler)

			So(err, ShouldBeNil)
			So(recv, ShouldNotBeNil)

			intent.Unregister(platform, recv)
		}))
	})
}

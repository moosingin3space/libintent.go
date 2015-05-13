package intent_test

import (
	"github.com/moosingin3space/libintent.go/intent"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"testing"
)

func WithMockData(f func(intent.Application, chan intent.Intent, func(intent.Intent) bool)) func() {
	return func() {
		app := intent.Application{
			Name:    "Awesome Listing Tool",
			Version: "0.0.1",
		}
		handler := make(chan intent.Intent)
		validator := func(intent intent.Intent) bool {
			return true
		}

		f(app, handler, validator)
	}
}

func WithPlatform(platform intent.Platform, f func(intent.Platform)) func() {
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
	Convey("Given a mock Platform", t, WithMockData(func(app intent.Application, handler chan intent.Intent, validator func(intent.Intent) bool) {
		Convey("Errors are passed up the stack", WithPlatform(FailsToMakeSocketPlatform{}, func(p intent.Platform) {
			recv, err := intent.Register(p, "ls", app, validator, handler)

			So(err, ShouldNotBeNil)
			So(recv, ShouldBeNil)
			So(err, ShouldEqual, io.ErrUnexpectedEOF)
		}))

		Convey("Unregistering a nil IntentReceiver has no effect", WithPlatform(FailsToMakeSocketPlatform{}, func(p intent.Platform) {
			var recv *intent.IntentReceiver = nil
			intent.Unregister(p, recv)
		}))
	}))
}

func TestDefaultPlatform(t *testing.T) {
	Convey("Given the default Platform", t, WithMockData(func(app intent.Application, handler chan intent.Intent, validator func(intent.Intent) bool) {
		Convey("Default platform initialization and destruction will not fail", WithPlatform(intent.DefaultPlatform(), nil))

		Convey("Intent registration and unregistration will not fail", WithPlatform(intent.DefaultPlatform(), func(platform intent.Platform) {
			recv, err := intent.Register(platform, "ls", app, validator, handler)

			So(err, ShouldBeNil)
			So(recv, ShouldNotBeNil)

			intent.Unregister(platform, recv)
		}))
	}))
}

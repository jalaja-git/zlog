package zlog

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogger(t *testing.T) {
	Convey("With a logger", t, func() {
		l := New("test_logger")
		Convey("I should be able to get name and level", func() {
			So(l.Name(), ShouldEqual, "test_logger")
			So(l.LogLevel(), ShouldEqual, Debug)
			So(l.Test(Debug), ShouldBeTrue)
		})

		Convey("I should be able to log a string with message", func() {
			b := &bytes.Buffer{}
			bl := l.Stream(b)
			bl.Info().Str("hello", "world").Msg("Test")
			So(strings.TrimSpace(b.String()), ShouldEqual, `{"level":"info","name":"test_logger","hello":"world","message":"Test"}`)
		})

		Convey("Disable Info log level", func() {
			wl := l.Level(Warn)
			So(wl.Info().Enabled(), ShouldBeFalse)
		})

		Convey("Fatal and panic events must be non-nil", func() {
			So(l.Fatal(), ShouldNotBeNil)
			So(l.Panic(), ShouldNotBeNil)
		})
	})
}

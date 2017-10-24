package zlog

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogger(t *testing.T) {
	Convey("Every logger must have a name and level", t, func() {
		l := New("test_logger")
		So(l.Name(), ShouldEqual, "test_logger")
		So(l.LogLevel(), ShouldEqual, Debug)
		So(l.Test(Debug), ShouldBeTrue)
		l.Info().Str("hello", "world").Msg("Test")
	})
}

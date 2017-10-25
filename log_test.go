package zlog

import (
	"errors"
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
		l.Info().Int("Int value", 1000).Msg("Printing Int")
		l.Info().Msg("Printing Message only")
		l.Info().Err(errors.New("Error event")).Msg("Printing Error event")

		log := l.With().Str("with_key", "with_val").Logger()
		log.Info().Msg("Str - With Testing")

		ilog := l.With().Int("with_key", 3000).Logger()
		ilog.Info().Msg("Int - With Testing")

		ilog = l.With().Int("with_key", 4000).Logger()
		ilog.Info().Msg("Int  2nd time - With Testing")

		ulog := ilog.With().Int("with_key", 5000).Logger()
		ulog.Info().Msg("Int  3rd time - With Testing")

		elog := l.With().Err(errors.New("some error")).Logger()
		elog.Info().Msg("Printing Err - With Testing")
		elog.Warn().Msg("Warning Printing Err - With Testing")

	})
}

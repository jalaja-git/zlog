package zlog_test

import (
	"errors"
	"fmt"
	stdlog "log"

	"github.com/rs/zlog"
)

func ExampleNew() {
	log := zlog.New("tst_logger")

	log.Info().Msg("hello world")

	// Output: {"level":"info","name":"tst_logger","message":"hello world"}
}

func ExampleLogger_Log() {
	log := zlog.New("tst_logger")

	log.Debug().
		Str("foo", "bar").
		Str("bar", "baz").
		Msg("")

	// Output: {"level":"debug","name":"tst_logger","foo":"bar","bar":"baz"}
}

func ExampleLogger_With() {
	log := zlog.New("tst_logger").
		With().
		Str("foo", "bar").
		Logger()

	log.Info().Msg("hello world")

	// Output: {"level":"info","name":"tst_logger","foo":"bar","message":"hello world"}
}
func ExampleLogger_Level() {
	log := zlog.New("tst_logger").Level(zlog.Warn)

	log.Info().Msg("filtered out message")
	log.Error().Msg("kept message")

	// Output: {"level":"error","name":"tst_logger","message":"kept message"}
}

func ExampleLogger_Info() {
	log := zlog.New("tst_logger")

	log.Info().
		Str("foo", "bar").
		Int("n", 123).
		Msg("hello world")

	// Output: {"level":"info","name":"tst_logger","foo":"bar","n":123,"message":"hello world"}
}

func ExampleLogger_Warn() {
	log := zlog.New("tst_logger")

	log.Warn().
		Str("foo", "bar").
		Msg("a warning message")

	// Output: {"level":"warn","name":"tst_logger","foo":"bar","message":"a warning message"}
}

func ExampleLogger_Error() {
	log := zlog.New("tst_logger")

	log.Error().
		Err(errors.New("some error")).
		Msg("error doing something")

	// Output: {"level":"error","name":"tst_logger","error":"some error","message":"error doing something"}
}

func ExampleLogger_Write() {
	log := zlog.New("tst_logger").With().
		Str("foo", "bar").
		Logger()

	stdlog.SetFlags(0)
	stdlog.SetOutput(log)

	stdlog.Print("hello world")

	// Output: {"name":"tst_logger","foo":"bar","message":"hello world"}
}

func ExampleEvent_Timestamp() {
	log := zlog.New("tst_logger")
	fmt.Println("h")

	log.Info().
		Timestamp().Msg("hello world")

	// Out: {"level":"info","name":"tst_logger","time":"2017-10-25T18:35:08-07:00","message":"hello world"}
}

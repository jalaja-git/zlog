package zlog_test

import (
	"errors"
	"fmt"
	stdlog "log"

	"github.com/atom-deps/zlog"
)

func ExampleNew() {
	log := zlog.New("tst_logger")

	log.Info().Msg("hello world")

	// Output: {"level":"info","name":"tst_logger","message":"hello world"}
}

func ExampleLogger_info() {
	log := zlog.New("tst_logger")

	log.Info().
		Str("foo", "bar").
		Int("n", 123).
		Bool("b", true).
		Msg("hello world")

	// Output: {"level":"info","name":"tst_logger","foo":"bar","n":123,"b":true,"message":"hello world"}
}

func ExampleLogger_warn() {
	log := zlog.New("tst_logger")

	log.Warn().
		Str("foo", "bar").
		Msg("a warning message")

	// Output: {"level":"warn","name":"tst_logger","foo":"bar","message":"a warning message"}
}

func ExampleLogger_error() {
	log := zlog.New("tst_logger")

	log.Error().
		Error(errors.New("some error")).
		Msg("error doing something")

	// Output: {"level":"error","name":"tst_logger","error":"some error","message":"error doing something"}
}
func ExampleLogger_msg() {
	log := zlog.New("tst_logger")

	log.Debug().
		Str("foo", "bar").
		Str("bar", "baz").
		Msg("")

	// Output: {"level":"debug","name":"tst_logger","foo":"bar","bar":"baz"}
}

func ExampleLogger_level() {
	log := zlog.New("tst_logger").Level(zlog.Warn)

	log.Info().Msg("filtered out message")
	log.Error().Msg("kept message")

	// Output: {"level":"error","name":"tst_logger","message":"kept message"}
}

func ExampleLogger_with() {
	log := zlog.New("tst_logger").
		With().
		Str("foo", "bar").
		Logger()

	log.Info().Msg("hello world")

	// Output: {"level":"info","name":"tst_logger","foo":"bar","message":"hello world"}
}

func ExampleLogger_write() {
	log := zlog.New("tst_logger").With().
		Str("foo", "bar").
		Logger()

	stdlog.SetFlags(0)
	stdlog.SetOutput(log)

	stdlog.Print("hello world")

	// Output: {"name":"tst_logger","foo":"bar","message":"hello world"}
}

func ExampleEvent_timestamp() {
	log := zlog.New("tst_logger")
	log = log.Level(zlog.Warn)
	// We cant test timestamp as its a moving target, lets filter it out
	// and figure it out later
	log.Info().Timestamp().Msg("hello world")
	log.Warn().Msg("hello world")

	// Output: {"level":"warn","name":"tst_logger","message":"hello world"}
}

func ExampleEvent_uint() {
	log := zlog.New("tst_logger")
	log.Info().Uint("uint", 1000).Msg("hello world")
	// Output: {"level":"info","name":"tst_logger","uint":1000,"message":"hello world"}
}

func ExampleEvent_float32() {
	log := zlog.New("tst_logger")
	log.Info().Float32("float32", 5.5555).Msg("hello world")
	// Output: {"level":"info","name":"tst_logger","float32":5.5555,"message":"hello world"}
}

func ExampleContext_with() {
	log := zlog.New("tst_logger")
	ctx := log.With()
	// Add Int, Bool and Str context to every log
	l := ctx.Int("n", 100).Bool("b", false).Str("t", "test").Uint("u", 1).Float32("f", 1.1).Logger()
	// Add Error context in addition to  above contexts.
	l = l.With().Error(nil).Logger()
	// Add one more Error Event in addition to above contexts.
	l = l.With().Error(fmt.Errorf("test error")).Logger()
	l.Info().Msg("context test")

	// Output: {"level":"info","name":"tst_logger","n":100,"b":false,"t":"test","u":1,"f":1.1,"error":"test error","message":"context test"}
}

type User struct {
	Name string
	ID   string
}

func (u *User) Marshal(e zlog.Event) {
	e.Str("name", u.Name).Str("id", u.ID)
}

func ExampleEvent_object() {
	log := zlog.New("tst_logger")
	log.Info().Object("user", &User{"Jalaja", "Ganapathy"}).Msg("object test")

	// Output: {"level":"info","name":"tst_logger","user":{"name":"Jalaja","id":"Ganapathy"},"message":"object test"}
}

func ExampleContext_with_object() {
	log := zlog.New("tst_logger")
	l := log.With().Object("user", &User{"Jalaja", "Ganapathy"}).Logger()
	l.Info().Msg("object test")

	// Output: {"level":"info","name":"tst_logger","user":{"name":"Jalaja","id":"Ganapathy"},"message":"object test"}
}

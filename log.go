package zlog

import (
	"io"
	"os"
)

// Level defines supported log levels.
type Level uint8

const (
	// Debug log level.
	Debug Level = iota

	// Info log level.
	Info

	// Warn log level.
	Warn

	// Error log level.
	Error

	// Fatal log level.
	Fatal

	// Panic log level.
	Panic

	// Disabled log level.
	Disabled
)

// Logger provides a leveled logging interface.
type Logger interface {
	// Name returns the name of the logger.
	Name() string

	// LogLevel returns the current level of the logger.
	LogLevel() Level

	// Test returns True if the logger will emit a log event at the given level.
	Test(lvl Level) bool

	// Write provides the io.Writer interface for this logger to be chained the standard
	// logging library.
	Write([]byte) (int, error)

	// Stream duplicates the current logger and sets the output to the given writer.
	Stream(w io.Writer) Logger

	// Level sets logger to the minimum accepted log level.
	Level(lvl Level) Logger

	// With creates a child logger with the field added to its context.
	With() Context

	// Debug starts a new message with debug level.
	// You must call Msg on the returned event in order to send the event.
	Debug() Event

	// Info starts a new message with info level.
	// You must call Msg on the returned event in order to send the event.
	Info() Event

	// Warn starts a new message with warn level.
	// You must call Msg on the returned event in order to send the event.
	Warn() Event

	// Error starts a new message with error level.
	// You must call Msg on the returned event in order to send the event.
	Error() Event

	// Fatal starts a new message with fatal level. The os.Exit(1) function
	// is called by the Msg method.
	//
	// You must call Msg on the returned event in order to send the event.
	Fatal() Event

	// Panic starts a new message with panic level. The message is also sent
	// to the panic function.
	//
	// You must call Msg on the returned event in order to send the event.
	Panic() Event
}

// New creates a new named logger at Debug level with output written to stdout.
func New(name string) Logger {
	return NewWithStream(name, os.Stdout, Debug)
}

// NewWithStream creates a new named logger with the given output writer and log level.
//
// Each logging operation must issue only a single call to the Writer's write method.
// The writer object is assumed to be thread-safe.
func NewWithStream(name string, w io.Writer, lvl Level) Logger {
	return newZeroLogger(name, w, lvl)
}

// Context configures a new sub-logger with contextual fields.
type Context interface {
	// Logger returns the logger with the context previously set.
	Logger() Logger

	// Str adds the field key with val as a string to the logger context.
	Str(key string, val string) Context

	// Int adds the field key with i as a int to the logger context.
	Int(key string, val int) Context

	// Err adds the field "error" with err as a string to the logger context.
	Error(e error) Context

	// Bool adds the field key with val as a bool to the logger context.
	Bool(key string, val bool) Context
}

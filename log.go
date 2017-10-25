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
	With() Context
	Debug() Event
	Info() Event
	Warn() Event
	Error() Event
	Fatal() Event
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

type Context interface {
	Logger() Logger
	Str(key string, val string) Context
	Int(key string, val int) Context
	Err(e error) Context
	Bool(key string, val bool) Context
}

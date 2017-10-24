package zlog

import (
	"io"

	"github.com/rs/zerolog"
)

// Trivial implementation using zerolog
type logImpl struct {
	zerolog.Logger
	name string
	lvl  Level
}

func (l logImpl) Name() string {
	return l.name
}

func (l logImpl) LogLevel() Level {
	return l.lvl
}

func (l logImpl) Test(lvl Level) bool {
	return lvl >= l.lvl
}

func (l logImpl) Write(b []byte) (int, error) {
	return l.Logger.Write(b)
}

func (l logImpl) Stream(w io.Writer) Logger {
	return logImpl{l.Logger.Output(w), l.name, l.lvl}
}

func (l logImpl) Level(lvl Level) Logger {
	zlvl := zerolog.Level(lvl)
	return logImpl{l.Logger.Level(zlvl), l.name, lvl}
}

func (l logImpl) With() Context {
	return ctxImpl{l.Logger.With(), l}
}

func (l logImpl) Debug() Event {
	return eventImpl{l.Logger.Debug()}
}

func (l logImpl) Info() Event {
	return eventImpl{l.Logger.Info()}
}

func (l logImpl) Warn() Event {
	return eventImpl{l.Logger.Warn()}
}

func (l logImpl) Error() Event {
	return eventImpl{l.Logger.Error()}
}

func (l logImpl) Fatal() Event {
	return eventImpl{l.Logger.Fatal()}
}

func (l logImpl) Panic() Event {
	return eventImpl{l.Logger.Panic()}
}

func newZeroLogger(name string, w io.Writer, lvl Level) Logger {
	zl := zerolog.New(w).Level(zerolog.Level(lvl)).With().Str("name", name).Logger()
	return logImpl{zl, name, Debug}
}

type eventImpl struct {
	*zerolog.Event
}

func (e eventImpl) Str(key string, val string) Event {
	e.Event = e.Event.Str(key, val)
	return e
}

func (e eventImpl) Int(key string, val int) Event {
	e.Event = e.Event.Int(key, val)
	return e
}

func (e eventImpl) Enabled() bool {
	return e.Event.Enabled()
}

func (e eventImpl) Msg(msg string) {
	e.Event.Msg(msg)
}

type ctxImpl struct {
	zerolog.Context
	l logImpl
}

func (c ctxImpl) Logger() Logger {
	return c.l
}

func (c ctxImpl) Str(key string, val string) Context {
	c.Context.Str(key, val)
	return c
}

func (c ctxImpl) Int(key string, val int) Context {
	c.Context.Int(key, val)
	return c
}

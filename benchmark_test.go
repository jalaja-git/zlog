package zlog

import (
	"errors"
	"io/ioutil"
	"testing"
)

var (
	errExample  = errors.New("fail")
	fakeMessage = "Test logging, but use a somewhat realistic message length."
)

func BenchmarkLogEmpty(b *testing.B) {
	logger := NewWithStream("test", ioutil.Discard, Debug)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg("")
		}
	})
}

func BenchmarkDisabled(b *testing.B) {
	logger := NewWithStream("test", ioutil.Discard, Disabled)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkInfo(b *testing.B) {
	logger := NewWithStream("test", ioutil.Discard, Debug)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkContextFields(b *testing.B) {
	logger := NewWithStream("test", ioutil.Discard, Debug).With().
		Str("string", "four!").
		Timestamp().
		Int("int", 123).
		Float32("float", -2.203230293249593).
		Logger()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkContextAppend(b *testing.B) {
	logger := NewWithStream("test", ioutil.Discard, Debug).With().
		Str("foo", "bar").
		Logger()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.With().Str("bar", "baz")
		}
	})
}

func BenchmarkLogFields(b *testing.B) {
	logger := NewWithStream("test", ioutil.Discard, Debug)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("string", "four!").
				Timestamp().
				Int("int", 123).
				Float32("float", -2.203230293249593).
				Msg(fakeMessage)
		}
	})
}

type obj struct {
	Pub  string
	Tag  string `json:"tag"`
	priv int
}

func (o obj) Marshal(e Event) {
	e.Str("Pub", o.Pub).
		Str("Tag", o.Tag).
		Int("priv", o.priv)
}

func BenchmarkLogFieldType(b *testing.B) {
	bools := []bool{true, false, true, false, true, false, true, false, true, false}
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	strings := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	objects := []obj{
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
		{"a", "a", 0},
	}
	errs := []error{errors.New("a"), errors.New("b"), errors.New("c"), errors.New("d"), errors.New("e")}
	types := map[string]func(e Event) Event{
		"Bool": func(e Event) Event {
			return e.Bool("k", bools[0])
		},
		"Int": func(e Event) Event {
			return e.Int("k", ints[0])
		},
		"Str": func(e Event) Event {
			return e.Str("k", strings[0])
		},
		"Err": func(e Event) Event {
			return e.Error(errs[0])
		},
		"Object": func(e Event) Event {
			return e.Object("k", objects[0])
		},
	}
	logger := NewWithStream("test", ioutil.Discard, Debug)
	b.ResetTimer()
	for name := range types {
		f := types[name]
		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					f(logger.Info()).Msg("")
				}
			})
		})
	}
}

# :mag: ZLOG - Structured Logging Interface

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]

zlog aims to provide a consistent, easy to use structured logging interface.

The interface is inspired by [zerolog](https://github.com/rs/zerolog) and uses
it for the default implementation that provides JSON.

zlog's API mimics the chaining API style from zerolog and logrus to provide a simple
intuitive interface to the developers while allowing high performance implementations
avoiding reflection and allocations.

Refer complete documentation : https://godoc.org/github.com/anuvu/zlog

## Features

* Context based logging
* Log level support
* Near-zero allocations

## Usage

```go
import "github.com/anuvu/zlog"
log.Info().Msg("hello world")

// Output: {"level":"info","time":1494567715,"message":"hello world"}
```

### Fields can be added to log messages

```go
log.Info().
    Str("foo", "bar").
    Int("n", 123).
    Msg("hello world")

// Output: {"level":"info",foo":"bar","n":123,"message":"hello world"}
```

### Create logger instance to manage different outputs

```go
logger := zlog.New("test_logger").With().Timestamp().Logger()

logger.Info().Str("foo", "bar").Msg("hello world")

// Output: {"level":"info","time":1494567715,"message":"hello world","foo":"bar"}
```

### Sub-loggers let you chain loggers with additional context

```go
sublogger := log.With().
                 Str("component": "foo").
                 Logger()
sublogger.Info().Msg("hello world")

// Output: {"level":"info","time":1494567715,"message":"hello world","component":"foo"}
```

### Set as standard logger output

```go
log := zlog.New("test_logger).With().
    Str("foo", "bar").
    Logger()

stdlog.SetFlags(0)
stdlog.SetOutput(log)

stdlog.Print("hello world")

// Output: {"foo":"bar","message":"hello world"}
```

## Field Types

### Standard Types

* `Str`
* `Bool`
* `Int` 
* `Uint` 
* `Float32`

### Advanced Fields

* `Error`: Takes an `error` and render it as a string using the `zlog.ErrorFieldName` field name.
* `Timestamp`: Insert a timestamp field with `zlog.TimestampFieldName` field name and formatted using `zlog.TimeFieldFormat`.

## Performance

```bash
$ go test -bench=. -benchmem
.........
9 total assertions

goos: darwin
goarch: amd64
pkg: github.com/anuvu/zlog
BenchmarkLogEmpty-8             50000000                31.2 ns/op             0 B/op          0 allocs/op
BenchmarkDisabled-8             200000000                6.37 ns/op            0 B/op          0 allocs/op
BenchmarkInfo-8                 30000000                50.7 ns/op             0 B/op          0 allocs/op
BenchmarkContextFields-8        10000000               143 ns/op               0 B/op          0 allocs/op
BenchmarkContextAppend-8        10000000               150 ns/op             832 B/op          3 allocs/op
BenchmarkLogFields-8            10000000               229 ns/op               0 B/op          0 allocs/op
BenchmarkLogFieldType/Int-8     30000000                49.1 ns/op             0 B/op          0 allocs/op
BenchmarkLogFieldType/Str-8     30000000                43.7 ns/op             0 B/op          0 allocs/op
BenchmarkLogFieldType/Err-8     30000000                49.5 ns/op             0 B/op          0 allocs/op
BenchmarkLogFieldType/Object-8  10000000               109 ns/op              64 B/op          2 allocs/op
BenchmarkLogFieldType/Bool-8    30000000                43.2 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/anuvu/zlog       17.846s
```

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://godoc.org/github.com/anuvu/zlog

[ci-img]: https://img.shields.io/travis/anuvu/zlog/master.svg
[ci]: https://travis-ci.org/uber-go/dig/branches

[cov-img]: https://codecov.io/gh/anuvu/zlog/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/anuvu/zlog/branch/master

[report-card-img]: https://goreportcard.com/badge/github.com/anuvu/zlog
[report-card]: https://goreportcard.com/report/github.com/anuvu/zlog

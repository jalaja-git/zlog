# :mag: ZLOG - Structured Logging Interface

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]

zlog aims to provide a consistent, easy to use structured logging interface.

The interface is inspired by [zerolog](https://github.com/rs/zerolog) and uses
it for the default implementation that provides JSON.

zlog's API mimics the chaining API style from zerolog and logrus to provide a simple
intuitive interface to the developers while allowing high performance implementations
avoiding reflection and allocations.

Refer complete documentation : https://godoc.org/github.com/atom-deps/zlog

## Features

* Context based logging
* Log level support
* Near-zero allocations

## Usage

```go
import "github.com/atom-deps/zlog"
log.Info().Msg("hello world")

// Output: {"level":"info","time":1494567715,"message":"hello world"}
```

### Fields can be added to log messages

```go
log.Info().
    Str("foo", "bar").
    Int("n", 123).
    Msg("hello world")

// Output: {"level":"info","time":1494567715,"foo":"bar","n":123,"message":"hello world"}
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


[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://godoc.org/github.com/atom-deps/zlog

[ci-img]: https://img.shields.io/travis/atom-deps/zlog/master.svg
[ci]: https://travis-ci.org/uber-go/dig/branches

[cov-img]: https://codecov.io/gh/atom-deps/zlog/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/atom-deps/zlog/branch/master

[report-card-img]: https://goreportcard.com/badge/github.com/atom-deps/zlog
[report-card]: https://goreportcard.com/report/github.com/atom-deps/zlog

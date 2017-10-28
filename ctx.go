package zlog

// Context configures a new sub-logger with contextual fields.
type Context interface {
	// Logger returns the logger with the context previously set.
	Logger() Logger

	// Str adds the field key with val as a string to the logger context.
	Str(key string, val string) Context

	// Int adds the field key with i as a int to the logger context.
	Int(key string, val int) Context

	// Uint adds the field key with i as a uint to the event.
	Uint(key string, i uint) Context

	// Float32 adds the field key with f as a float32 to the event.
	Float32(key string, f float32) Context

	// Bool adds the field key with val as a bool to event.
	Bool(key string, val bool) Context

	// Err adds the field "error" with err as a string to the event.
	// If err is nil, no field is added.
	Error(e error) Context

	// Timestamp adds the current local time as UNIX timestamp to the event.
	Timestamp() Context

	// Object marshals a custom type that implements ObjectMarshaler interface.
	Object(key string, obj ObjectMarshaler) Context
}

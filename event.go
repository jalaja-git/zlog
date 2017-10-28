package zlog

// ObjectMarshaler provides a strongly-typed encoding agnostic interface
// to be implemented by custom types for efficient encoding
type ObjectMarshaler interface {
	Marshal(e Event)
}

// Event provides a log event interface.
type Event interface {

	// Enabled return false if the *Event is going to be filtered out by
	// log level or sampling.
	Enabled() bool

	// Msg sends the Event with msg added as the message field if not empty.
	//
	// NOTICE: once this method is called, the Event should be disposed.
	// Calling Msg twice can have unexpected result.
	Msg(msg string)

	// Str adds the field key with val as a string to the event.
	Str(key string, val string) Event

	// Int adds the field key with i as a int to the event.
	Int(key string, val int) Event

	// Uint adds the field key with i as a uint to the event.
	Uint(key string, i uint) Event

	// Float32 adds the field key with f as a float32 to the event.
	Float32(key string, f float32) Event

	// Bool adds the field key with val as a bool to event.
	Bool(key string, val bool) Event

	// Err adds the field "error" with err as a string to the event.
	// If err is nil, no field is added.
	Error(e error) Event

	// Timestamp adds the current local time as UNIX timestamp to the event.
	Timestamp() Event

	// Object marshals a custom type that implements ObjectMarshaler interface.
	Object(key string, obj ObjectMarshaler) Event
}

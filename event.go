package zlog

type Event interface {

	// Enabled return false if the *Event is going to be filtered out by
	// log level or sampling.
	Enabled() bool

	// Msg sends the Event with msg added as the message field if not empty.
	//
	// NOTICE: once this method is called, the Event should be disposed.
	// Calling Msg twice can have unexpected result.
	Msg(msg string)

	// Str adds the field key with val as a string to the Event context.
	Str(key string, val string) Event

	// Int adds the field key with i as a int to the Event context.
	Int(key string, val int) Event

	// Err adds the field "error" with err as a string to the Event context.
	// If err is nil, no field is added.
	Error(e error) Event

	// Timestamp adds the current local time as UNIX timestamp to the Event context with the "time" key.
	Timestamp() Event
}

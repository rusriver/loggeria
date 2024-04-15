package loggeria

import (
	"io"
)

type ILogger interface {
	// Event initializers
	DBG(func(e E)) IEvent
	DIS(func(e E)) IEvent
	IfErr(err error, f func(e E)) IEvent // automatically fizzles (does nothing), if err == nil
	ERR(func(e E)) IEvent
	FTL(func(e E)) IEvent
	INF(func(e E)) IEvent
	PNC(func(e E)) IEvent
	TRC(func(e E)) IEvent
	WRN(func(e E)) IEvent
	LEVEL(Level, func(e E)) IEvent

	// Sub-system initializer; to be used right before event initializer
	Sub(s string) ILogger

	// construct nested data
	NewMap() IMapBuffer
	// NewList() IListBuffer

	// Logger management
	// All functions returning ILogger, do this for convenience only, the actual effect must be
	// applied to original object as well; i.e. the logger object MUST be passed by pointer, not by value.
	// The only exception is Derive(), which explicitly creates a copy.
	AutoTimestamp() ILogger           // Add a "time" key automatically to each event
	CallerSkipFrame(skip int) ILogger // CallerSkipFrame instructs any future Caller calls to skip the specified number of frames. This includes those added via hooks from the context. The parameter will be set on IMapBuffer
	Derive() ILogger                  // make a derivative logger, with buffer copied; the same as With() in original subzerologger
	GetBuffer() IMapBuffer            // Can be used to set Mapped Diagnostic Context (default data for whole logger).
	Hook(hooks ...IHook) ILogger
	SetLevel(lvl Level) ILogger
	GetLevel() Level
	SetOutputWriter(w io.Writer) ILogger
	SetSampler(s ISampler) ILogger
	Write(p []byte) (n int, err error) // Write implements the io.Writer interface. This is useful to set as a writer for the standard library log.
}

// shorter alias
type E IMapBuffer

type IEvent interface {
	Send()
}

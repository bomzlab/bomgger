/*
	bomz log interface

	@since 0.1
	@version 0.1
*/
package log

import ()

// bomz logger interface
type Logger interface {
	bomzTrace
	bomzDebug
	bomzInfo
	bomzWarn
	bomzError
	bomzPanic
}

// bomz trace interface
type bomzTrace interface {
	trace(format string, v ...interface{})
	IsTrace() bool
}

// bomz debug interface
type bomzDebug interface {
	Debug(format string, v ...interface{})
	IsDebug() bool
}

// bomz info interface
type bomzInfo interface {
	Info(format string, v ...interface{})
	IsInfo() bool
}

// bomz warn interface
type bomzWarn interface {
	Warn(format string, v ...interface{})
	IsWarn() bool
}

// bomz error interface
type bomzError interface {
	Error(format string, v ...interface{})
	IsError() bool
}

// bomz panic interface
type bomzPanic interface {
	Panic(format string, v ...interface{})
	IsPanic() bool
}

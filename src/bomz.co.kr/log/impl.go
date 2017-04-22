/*
	bomz log impl

	implements bomz.co.kr/log/Logger interface

	@since 0.1
	@version 0.1
*/
package log

import ()

// log level
const(
	trace iota
	debug 
	info
	warn
	err
	pan
)

type impl struct {
	levels [6]bool
	logs []Logger
}

func (i *impl) Trace(format string, v ...interface{}){
	
}

func (i *impl) IsTrace() bool{
	return i.levels[trace];
}

func (i *impl) Debug(format string, v ...interface{}){
	
}

func (i *impl) IsDebug() bool{
	return i.levels[debug];
}

func (i *impl) Info(format string, v ...interface{}){
	
}

func (i *impl) IsInfo() bool{
	return i.levels[info];
}

func (i *impl) Warn(format string, v ...interface{}){
	
}

func (i *impl) IsWarn() bool{
	return i.levels[warn];
}

func (i *impl) Error(format string, v ...interface{}){
	
}

func (i *impl) IsError() bool{
	return i.levels[error];
}

func (i *impl) Panic(format string, v ...interface{}){
	
}

func (i *impl) IsPanic() bool{
	return i.levels[pan];
}
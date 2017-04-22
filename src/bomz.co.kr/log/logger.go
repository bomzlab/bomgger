/*
	bomz log

	@since 0.1
	@version 0.1
*/
package log

import (
	"fmt"
)

type log struct {
	loggerMap map[string]Logger
}

// singleton pattern log
var log = &log{}

var initFile string = "ddd.ymul"

func SetInitFile(nm string) {
	initFile = nm
}

/*
	get log instance
	if nm is nil or unknow name that return console logger
*/
func GetInstance(nm string) Logger {
	if nm == nil {
		// default console logger
		return consoleLogger()
	}

	if result := log[nm]; result != nil {
		return result
	}

	// unknow nm is default console logger return
	return consoleLogger()
}

// default console logger
func consolLogger() Logger {
	return nil
}

package agent

import "log"

var Verbose bool

// TODO: using https://github.com/golang/glog for logging
func Info(format string, v ...interface{}) {
	if Verbose {
		log.Printf(format, v...)
	}
}

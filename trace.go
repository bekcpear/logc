//go:build !notracelog

package logc

import "fmt"

// Trace print the message in the trace level
func Trace(msg ...any) {
	defaultLogC.Trace(msg...)
}

// Tracef act like Trace but with format
func Tracef(format string, msg ...any) {
	defaultLogC.Tracef(format, msg...)
}

// Trace print the message in the trace level
func (lc *LogC) Trace(msg ...any) {
	if lc.conf.level >= LevelTrace {
		lc.output(TraceLog, fmt.Sprintln(msg...))
	}
}

// Tracef act like Trace but with format
func (lc *LogC) Tracef(format string, msg ...any) {
	if lc.conf.level >= LevelTrace {
		lc.output(TraceLog, fmt.Sprintf(format, msg...))
	}
}

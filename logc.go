package logc

import (
	"log"
	"os"
)

var traceLog = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
var stdoutLog = log.New(os.Stdout, "", log.LstdFlags)
var stderrLog = log.New(os.Stderr, "", log.LstdFlags)

// SetPrefix is used to set the prefix of the log
func SetPrefix(prefix string) {
	traceLog.SetPrefix(prefix)
	stdoutLog.SetPrefix(prefix)
	stderrLog.SetPrefix(prefix)
}

// Trace print the message in the trace level
func Trace(msg any) {
	if myConf.Level <= LevelTrace {
		traceLog.Println(msg)
	}
}

// Debug print the message in the debug level
func Debug(msg any) {
	if myConf.Level <= LevelDebug {
		stdoutLog.Println(msg)
	}
}

// Info print the message in the info level
func Info(msg any) {
	if myConf.Level <= LevelInfo {
		stdoutLog.Println(msg)
	}
}

// Log print the message in the Normal level
func Log(msg any) {
	if myConf.Level <= LevelNormal {
		stdoutLog.Println(msg)
	}
}

// Warn print the message in the warning level
func Warn(msg any) {
	if myConf.Level <= LevelWarning {
		stderrLog.Println(msg)
	}
}

// Error print the message in the error level
func Error(msg any) {
	if myConf.Level <= LevelError {
		stderrLog.Println(msg)
	}
}

// Fatal print the message in the fatal level, and exit the program
func Fatal(msg any) {
	if myConf.Level <= LevelFatal {
		stderrLog.Fatalln(msg)
	}
}

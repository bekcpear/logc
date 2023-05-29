package logc

import (
	"fmt"
	"log"
	"os"
)

const (
	tracePrefix  = "TRACE"
	debugPrefix  = "DEBUG"
	infoPrefix   = "INFO"
	normalPrefix = "NORMAL"
	warnPrefix   = "WARN"
	errorPrefix  = "ERROR"
	fatalPrefix  = "FATAL"
)

var (
	traceLog  = log.New(os.Stdout, tracePrefix, log.LstdFlags|log.Llongfile)
	debugLog  = log.New(os.Stdout, debugPrefix, log.LstdFlags)
	infoLog   = log.New(os.Stdout, infoPrefix, log.LstdFlags)
	normalLog = log.New(os.Stdout, normalPrefix, log.LstdFlags)
	warnLog   = log.New(os.Stderr, warnPrefix, log.LstdFlags)
	errorLog  = log.New(os.Stderr, errorPrefix, log.LstdFlags)
	fatalLog  = log.New(os.Stderr, fatalPrefix, log.LstdFlags)
)

// SetPrefix is used to set the prefix of the log
func SetPrefix(prefix string) {
	traceLog.SetPrefix(fmt.Sprintf("%s %s", tracePrefix, prefix))
	debugLog.SetPrefix(fmt.Sprintf("%s %s", debugPrefix, prefix))
	infoLog.SetPrefix(fmt.Sprintf("%s %s", infoPrefix, prefix))
	normalLog.SetPrefix(fmt.Sprintf("%s %s", normalPrefix, prefix))
	warnLog.SetPrefix(fmt.Sprintf("%s %s", warnPrefix, prefix))
	errorLog.SetPrefix(fmt.Sprintf("%s %s", errorPrefix, prefix))
	fatalLog.SetPrefix(fmt.Sprintf("%s %s", fatalPrefix, prefix))
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
		debugLog.Println(msg)
	}
}

// Info print the message in the info level
func Info(msg any) {
	if myConf.Level <= LevelInfo {
		infoLog.Println(msg)
	}
}

// Log print the message in the Normal level
func Log(msg any) {
	if myConf.Level <= LevelNormal {
		normalLog.Println(msg)
	}
}

// Warn print the message in the warning level
func Warn(msg any) {
	if myConf.Level <= LevelWarning {
		warnLog.Println(msg)
	}
}

// Error print the message in the error level
func Error(msg any) {
	if myConf.Level <= LevelError {
		errorLog.Println(msg)
	}
}

// Fatal print the message in the fatal level, and exit the program
func Fatal(msg any) {
	if myConf.Level <= LevelFatal {
		fatalLog.Fatalln(msg)
	}
}

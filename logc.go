package logc

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

const (
	tracePrefix  = "(T) "
	debugPrefix  = "(D) "
	infoPrefix   = "(I) "
	normalPrefix = "(N) "
	warnPrefix   = "(W) "
	errorPrefix  = "(E) "
	fatalPrefix  = "(F) "
)

type LogC struct {
	traceLog  *log.Logger
	debugLog  *log.Logger
	infoLog   *log.Logger
	normalLog *log.Logger
	warnLog   *log.Logger
	errorLog  *log.Logger
	fatalLog  *log.Logger
	conf      *Configuration
}

func NewLogC() *LogC {
	return &LogC{
		traceLog: log.New(os.Stdout, tracePrefix,
			log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds),
		debugLog: log.New(os.Stdout, debugPrefix,
			log.LstdFlags|log.Lmsgprefix),
		infoLog: log.New(os.Stdout, infoPrefix,
			log.LstdFlags|log.Lmsgprefix),
		normalLog: log.New(os.Stdout, normalPrefix,
			log.LstdFlags|log.Lmsgprefix),
		warnLog: log.New(os.Stderr, warnPrefix,
			log.LstdFlags|log.Lmsgprefix),
		errorLog: log.New(os.Stderr, errorPrefix,
			log.LstdFlags|log.Lmsgprefix),
		fatalLog: log.New(os.Stderr, fatalPrefix,
			log.LstdFlags|log.Lmsgprefix),
		conf: &Configuration{HideToken: true},
	}
}

var defaultLogC = NewLogC()

func Default() *LogC {
	return defaultLogC
}

// SetPrefix is used to append a unique additional prefix to the current
// prefix.
func SetPrefix(prefix string) {
	defaultLogC.SetPrefix(prefix)
}

// Debug print the message in the debug level
func Debug(msg ...any) {
	defaultLogC.Debug(msg...)
}

// Info print the message in the info level
func Info(msg ...any) {
	defaultLogC.Info(msg...)
}

// Log print the message in the Normal level
func Log(msg ...any) {
	defaultLogC.Log(msg...)
}

// Warn print the message in the warning level
func Warn(msg ...any) {
	defaultLogC.Warn(msg...)
}

// Error print the message in the error level
func Error(msg ...any) {
	defaultLogC.Error(msg...)
}

// Fatal print the message in the fatal level, and exit the program
func Fatal(msg ...any) {
	defaultLogC.Fatal(msg...)
}

// Debugf act like Debug but with format
func Debugf(format string, msg ...any) {
	defaultLogC.Debugf(format, msg...)
}

// Infof act like Info but with format
func Infof(format string, msg ...any) {
	defaultLogC.Infof(format, msg...)
}

// Logf act like Log but with format
func Logf(format string, msg ...any) {
	defaultLogC.Logf(format, msg...)
}

// Warnf act like Warn but with format
func Warnf(format string, msg ...any) {
	defaultLogC.Warnf(format, msg...)
}

// Errorf act like Error but with format
func Errorf(format string, msg ...any) {
	defaultLogC.Errorf(format, msg...)
}

// Fatalf act like Fatal but with format
func Fatalf(format string, msg ...any) {
	defaultLogC.Fatalf(format, msg...)
}

// SetPrefix is used to append a unique additional prefix to the current
// prefix.
func (lc *LogC) SetPrefix(prefix string) {
	lc.traceLog.SetPrefix(fmt.Sprintf("%s%s ", tracePrefix, prefix))
	lc.debugLog.SetPrefix(fmt.Sprintf("%s%s ", debugPrefix, prefix))
	lc.infoLog.SetPrefix(fmt.Sprintf("%s%s ", infoPrefix, prefix))
	lc.normalLog.SetPrefix(fmt.Sprintf("%s%s ", normalPrefix, prefix))
	lc.warnLog.SetPrefix(fmt.Sprintf("%s%s ", warnPrefix, prefix))
	lc.errorLog.SetPrefix(fmt.Sprintf("%s%s ", errorPrefix, prefix))
	lc.fatalLog.SetPrefix(fmt.Sprintf("%s%s ", fatalPrefix, prefix))
}

// Debug print the message in the debug level
func (lc *LogC) Debug(msg ...any) {
	if lc.conf.Level >= LevelDebug {
		lc.debugLog.Println(msg...)
	}
}

// Info print the message in the info level
func (lc *LogC) Info(msg ...any) {
	if lc.conf.Level >= LevelInfo {
		lc.infoLog.Println(msg...)
	}
}

// Log print the message in the Normal level
func (lc *LogC) Log(msg ...any) {
	if lc.conf.Level >= LevelNormal {
		lc.normalLog.Println(msg...)
	}
}

// Warn print the message in the warning level
func (lc *LogC) Warn(msg ...any) {
	if lc.conf.Level >= LevelWarning {
		lc.warnLog.Println(msg...)
	}
}

// Error print the message in the error level
func (lc *LogC) Error(msg ...any) {
	if lc.conf.Level >= LevelError {
		lc.errorLog.Println(msg...)
	}
}

// Fatal print the message in the fatal level, and exit the program
func (lc *LogC) Fatal(msg ...any) {
	if lc.conf.Level >= LevelFatal {
		newItems := make([]any, 1)

		skip := 1
		_, file, num, ok := runtime.Caller(skip)
		fa := strings.Split(file, "/")
		for fa[len(fa)-3] == "oss.ac" &&
			fa[len(fa)-2] == "logc" &&
			strings.Split(fa[len(fa)-1], ":")[0] == "logc.go" {
			skip++
			_, file, num, ok = runtime.Caller(skip)
			fa = strings.Split(file, "/")
		}
		if ok {
			newItems[0] = fmt.Sprintf("%s:%d:", file, num)
		}
		newItems = append(newItems, msg...)
		lc.fatalLog.Fatalln(newItems...)
	}
}

// Debugf act like Debug but with format
func (lc *LogC) Debugf(format string, msg ...any) {
	if lc.conf.Level >= LevelDebug {
		lc.debugLog.Printf(format, msg...)
	}
}

// Infof act like Info but with format
func (lc *LogC) Infof(format string, msg ...any) {
	if lc.conf.Level >= LevelInfo {
		lc.infoLog.Printf(format, msg...)
	}
}

// Logf act like Log but with format
func (lc *LogC) Logf(format string, msg ...any) {
	if lc.conf.Level >= LevelNormal {
		lc.normalLog.Printf(format, msg...)
	}
}

// Warnf act like Warn but with format
func (lc *LogC) Warnf(format string, msg ...any) {
	if lc.conf.Level >= LevelWarning {
		lc.warnLog.Printf(format, msg...)
	}
}

// Errorf act like Error but with format
func (lc *LogC) Errorf(format string, msg ...any) {
	if lc.conf.Level >= LevelError {
		lc.errorLog.Printf(format, msg...)
	}
}

// Fatalf act like Fatal but with format
func (lc *LogC) Fatalf(format string, msg ...any) {
	if lc.conf.Level >= LevelFatal {
		skip := 1
		_, file, num, ok := runtime.Caller(skip)
		fa := strings.Split(file, "/")
		for fa[len(fa)-3] == "oss.ac" &&
			fa[len(fa)-2] == "logc" &&
			strings.Split(fa[len(fa)-1], ":")[0] == "logc.go" {
			skip++
			_, file, num, ok = runtime.Caller(skip)
			fa = strings.Split(file, "/")
		}
		var p string
		if ok {
			p = fmt.Sprintf("%s:%d:", file, num)
		}
		o := fmt.Sprintf(format, msg...)
		lc.fatalLog.Fatalln(p, o)
	}
}

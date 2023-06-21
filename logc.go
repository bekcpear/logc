package logc

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"oss.ac/hidapp/pkg/hidapp"
)

const (
	TraceLog = iota
	DebugLog
	InfoLog
	NormalLog
	WarnLog
	ErrorLog
	FatalLog
	loggersCount
)

var commonPrefix = [loggersCount]string{
	"(T) ",
	"(D) ",
	"(I) ",
	"(N) ",
	"(W) ",
	"(E) ",
	"(F) ",
}

type LogC struct {
	loggers [loggersCount]*log.Logger
	conf    *configuration
	hidePP  *hidapp.Processor
}

func NewLogC() *LogC {
	lc := new(LogC)
	dw := os.Stdout
	lc.loggers[TraceLog] = log.New(os.Stdout, commonPrefix[TraceLog],
		log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds)
	for i := 1; i < loggersCount; i++ {
		if i >= WarnLog {
			dw = os.Stderr
		}
		lc.loggers[i] = log.New(dw, commonPrefix[i],
			log.LstdFlags|log.Lmsgprefix)
	}

	lc.hidePP = hidapp.NewProcessor()
	lc.conf = &configuration{hideToken: true}
	return lc
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

// SetOutput set the output of loggers in the default LogC instance
// specified by index, the index is optional, if no one provided,
// the all loggers will set the output to w.
func SetOutput(w io.Writer, index ...int) {
	defaultLogC.SetOutput(w, index...)
}

// StdoutTo set the default stdout loggers to new writer.
func StdoutTo(w io.Writer) {
	defaultLogC.StdoutTo(w)
}

// StderrTo set the default stderr loggers to new writer.
func StderrTo(w io.Writer) {
	defaultLogC.StderrTo(w)
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
	for i := 0; i < loggersCount; i++ {
		lc.loggers[i].SetPrefix(fmt.Sprintf("%s%s ", commonPrefix[i], prefix))
	}
}

// SetOutput set the output of loggers in this LogC instance specified by
// index, the index is optional, if no one provided,
// the all loggers will set the output to w.
func (lc *LogC) SetOutput(w io.Writer, index ...int) {
	if len(index) == 0 {
		for i := 0; i < loggersCount; i++ {
			lc.loggers[i].SetOutput(w)
		}
	} else {
		for i := 0; i < len(index); i++ {
			lc.loggers[index[i]].SetOutput(w)
		}
	}
}

// StdoutTo set the stdout loggers to the new writer.
func (lc *LogC) StdoutTo(w io.Writer) {
	for i := 0; i < WarnLog; i++ {
		v, ok := lc.loggers[i].Writer().(*os.File)
		if ok {
			if v.Name() == "/dev/stdout" {
				lc.loggers[i].SetOutput(w)
			}
		}
	}
}

// StderrTo set the stderr loggers to the new writer.
func (lc *LogC) StderrTo(w io.Writer) {
	for i := WarnLog; i < loggersCount; i++ {
		v, ok := lc.loggers[i].Writer().(*os.File)
		if ok {
			if v.Name() == "/dev/stderr" {
				lc.loggers[i].SetOutput(w)
			}
		}
	}
}

func (lc *LogC) output(index int, s string) {
	if lc.conf.hideToken {
		s = lc.hidePP.Process(s)
	}
	lc.loggers[index].Output(2, s)
}

// Debug print the message in the debug level
func (lc *LogC) Debug(msg ...any) {
	if lc.conf.level >= LevelDebug {
		lc.output(DebugLog, fmt.Sprintln(msg...))
	}
}

// Info print the message in the info level
func (lc *LogC) Info(msg ...any) {
	if lc.conf.level >= LevelInfo {
		lc.output(InfoLog, fmt.Sprintln(msg...))
	}
}

// Log print the message in the Normal level
func (lc *LogC) Log(msg ...any) {
	if lc.conf.level >= LevelNormal {
		lc.output(NormalLog, fmt.Sprintln(msg...))
	}
}

// Warn print the message in the warning level
func (lc *LogC) Warn(msg ...any) {
	if lc.conf.level >= LevelWarning {
		lc.output(WarnLog, fmt.Sprintln(msg...))
	}
}

// Error print the message in the error level
func (lc *LogC) Error(msg ...any) {
	if lc.conf.level >= LevelError {
		lc.output(ErrorLog, fmt.Sprintln(msg...))
	}
}

// Fatal print the message in the fatal level, and exit the program
func (lc *LogC) Fatal(msg ...any) {
	if lc.conf.level >= LevelFatal {
		newItems := make([]any, 1)

		skip := 1
		_, file, num, ok := runtime.Caller(skip)
		fa := strings.Split(file, "/")
		for fa[len(fa)-3] == "oss.ac" &&
			fa[len(fa)-2][:4] == "logc" &&
			strings.Split(fa[len(fa)-1], ":")[0] == "logc.go" {
			skip++
			_, file, num, ok = runtime.Caller(skip)
			fa = strings.Split(file, "/")
		}
		if ok {
			newItems[0] = fmt.Sprintf("%s:%d:", file, num)
		}
		newItems = append(newItems, msg...)
		lc.output(FatalLog, fmt.Sprintln(newItems...))
		os.Exit(1)
	}
}

// Debugf act like Debug but with format
func (lc *LogC) Debugf(format string, msg ...any) {
	if lc.conf.level >= LevelDebug {
		lc.output(DebugLog, fmt.Sprintf(format, msg...))
	}
}

// Infof act like Info but with format
func (lc *LogC) Infof(format string, msg ...any) {
	if lc.conf.level >= LevelInfo {
		lc.output(InfoLog, fmt.Sprintf(format, msg...))
	}
}

// Logf act like Log but with format
func (lc *LogC) Logf(format string, msg ...any) {
	if lc.conf.level >= LevelNormal {
		lc.output(NormalLog, fmt.Sprintf(format, msg...))
	}
}

// Warnf act like Warn but with format
func (lc *LogC) Warnf(format string, msg ...any) {
	if lc.conf.level >= LevelWarning {
		lc.output(WarnLog, fmt.Sprintf(format, msg...))
	}
}

// Errorf act like Error but with format
func (lc *LogC) Errorf(format string, msg ...any) {
	if lc.conf.level >= LevelError {
		lc.output(ErrorLog, fmt.Sprintf(format, msg...))
	}
}

// Fatalf act like Fatal but with format
func (lc *LogC) Fatalf(format string, msg ...any) {
	if lc.conf.level >= LevelFatal {
		skip := 1
		_, file, num, ok := runtime.Caller(skip)
		fa := strings.Split(file, "/")
		for fa[len(fa)-3] == "oss.ac" &&
			fa[len(fa)-2][:4] == "logc" &&
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
		lc.output(FatalLog, p+" "+o)
		os.Exit(1)
	}
}

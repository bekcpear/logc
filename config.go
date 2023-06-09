package logc

import "log"

const (
	LevelFatal = iota - 3
	LevelError
	LevelWarning
	LevelNormal
	LevelInfo
	LevelDebug
	LevelTrace
)

type Configuration struct {
	Level     int
	HideToken bool
}

// TODO: hide token

// SetLogLevel is used to set the log level,
// the default log level is LevelNormal
func SetLogLevel(level int) {
	defaultLogC.SetLogLevel(level)
}

// DontHideToken is used to un-hide the tokens within log message,
// the default behavior is hiding those tokens.
func DontHideToken() {
	defaultLogC.DontHideToken()
}

// SetLogLevel is used to set the log level,
// the default log level is LevelNormal
func (lc *LogC) SetLogLevel(level int) {
	if level < LevelFatal || level > LevelTrace {
		log.Fatalln("the log level should between -3 and 3")
	}
	lc.conf.Level = level
}

// DontHideToken is used to un-hide the tokens within log message,
// the default behavior is hiding those tokens.
func (lc *LogC) DontHideToken() {
	lc.conf.HideToken = false
}

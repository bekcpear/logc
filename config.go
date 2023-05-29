package logc

import "log"

const (
	LevelTrace = iota - 3
	LevelDebug
	LevelInfo
	LevelNormal
	LevelWarning
	LevelError
	LevelFatal
)

type Configuration struct {
	Level     int
	HideToken bool
}

var myConf = Configuration{HideToken: true}

// SetLogLevel is used to set the log level,
// the default log level is LevelNormal
func SetLogLevel(level int) {
	if level < LevelTrace || level > LevelFatal {
		log.Fatalln("the log level should between -3 and 3")
	}
	myConf.Level = level
}

// DontHideToken is used to un-hide the tokens within log message,
// the default behavior is hiding those tokens.
func DontHideToken() {
	myConf.HideToken = false
}

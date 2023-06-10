package logc

import (
	"log"
)

const (
	LevelFatal = iota - 3
	LevelError
	LevelWarning
	LevelNormal
	LevelInfo
	LevelDebug
	LevelTrace
)

type configuration struct {
	level     int
	hideToken bool
}

// SetLogLevel is used to set the log level,
// the default log level is LevelNormal
func SetLogLevel(level int) {
	defaultLogC.SetLogLevel(level)
}

// AppendPassphraseRegexp is used to append the regexps to match the
// tokens/passphrases you want to hide them in the log output.
// Here are two type supported, string and *regexp.Regexp.
func AppendPassphraseRegexp(re ...any) error {
	return defaultLogC.hidePP.AppendRegexp(re...)
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
	lc.conf.level = level
}

// AppendPassphraseRegexp is used to append the regexps to match the
// tokens/passphrases you want to hide them in the log output.
// Here are two type supported, string and *regexp.Regexp.
func (lc *LogC) AppendPassphraseRegexp(re ...any) error {
	return lc.hidePP.AppendRegexp(re...)
}

// DontHideToken is used to un-hide the tokens within log message,
// the default behavior is hiding those tokens.
func (lc *LogC) DontHideToken() {
	lc.conf.hideToken = false
}

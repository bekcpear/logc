//go:build notracelog

package logc

// Trace do nothing
func Trace(msg ...any) {
	return
}

// Tracef do nothing
func Tracef(format string, msg ...any) {
	return
}

// Trace print the message in the trace level
func (lc *LogC) Trace(msg ...any) {
	return
}

// Tracef act like Trace but with format
func (lc *LogC) Tracef(format string, msg ...any) {
	return
}

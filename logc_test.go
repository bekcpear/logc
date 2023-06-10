package logc

import (
	"reflect"
	"testing"
)

type cWriter struct {
	o int
}

func (cw *cWriter) Write(p []byte) (nw int, error error) {
	return
}

func TestLogC_StdXXXTo(t *testing.T) {
	lc := NewLogC()
	cw := new(cWriter)
	cw1 := new(cWriter)
	cw1.o = 2

	lc.SetOutput(cw1, DebugLog)

	lc.StdoutTo(cw)
	lc.StderrTo(cw)

	for i := 0; i < loggersCount; i++ {
		w := lc.loggers[i].Writer()
		if i == DebugLog {
			if !reflect.DeepEqual(cw1, w) {
				t.Fatalf("%d: unmatched cWriter: %#v <-> %#v\n", i, cw1, w)
			}
		} else {
			if !reflect.DeepEqual(cw, w) {
				t.Fatalf("%d: unmatched cWriter: %#v <-> %#v\n", i, cw, w)
			}
		}
	}
}

package log

import "testing"

func TestLog(t *testing.T) {
	Info("info")
	Warn("warn")
	Error("error")
}

func TestSetConfigStr(t *testing.T) {
	glogger := SetConfigString(`{ mike {formatter.name = "json" }`)
	glogger.Logger("mike").logger.Warn("hello json")
}

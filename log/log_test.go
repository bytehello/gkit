package log

import "testing"

func TestLog(t *testing.T) {
	Info("info")
	Warn("warn")
	Error("error")
}

func TestSetConfigStr(t *testing.T) {
	_, _ = SetConfigFile("mate.conf")
	SetReportCaller(true)
	Logger().Info("default log from Glogger")
	Logger("mike").Info("mike log from Glogger")

	Info("default log use hijack Info")
	Warn("default log use hijack Warn")
}

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
	Debug("default log debug")
	Trace("default log trace")
}

func TestSetLogger(t *testing.T) {
	_, err := SetConfigFile("mate.conf")
	t.Log(err)
	if SetLogger("mike") == nil {
		t.Error("set logger err")
	}
	Info("info")
}

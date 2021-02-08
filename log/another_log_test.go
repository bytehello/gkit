package log

import (
	//nlog "log"
	"testing"

	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
)

// 创建一个新的logger
// 使用场景：需要在特定的文件夹打印的情况
// 这里返回的是logrus.Logger 对象
// TODO: gkit可以对其封装，返回一个gkit对应的日志库对象

func TestLogMate(t *testing.T) {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString( // 这里使用的是一种特殊的数据格式，号称是更加容易读
			`{ mike {formatter.name = "json" } { default {formatter.name = "text"} }`,
		),
	)

	mate.Logger("mike").Errorln("With JSON MIKE")
	mate.Logger().Error("another error logger") // 不传入参数，默认是default
}

// 劫持使用，将logrus.Logger的行为替换掉，自己的项目中有这种情况没？
// TODO 有这种需求可以，封装
// mike 代表的配置劫持了StandardLogger
func TestHijactLoggerByMate(t *testing.T) {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
	)
	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)
	logrus.Println("hello std logger is hijack by mike")
}

// mike 代表的配置劫持了 原生 log
// NOTICE :这个是不被允许的，只能劫持logrus.Logger
func TestHijactNativeLoggerByMate(t *testing.T) {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
	)
	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)
	logrus.Println("hello std logger is hijack by mike")
}

func TestHijackLoggerOverwriteByMate(t *testing.T) {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
		// 读取文件覆盖configString
		logrus_mate.ConfigFile("mate.conf"),
	)
	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)
	logrus.Println("hello std logger is hijack by mike")
}

func TestLogrusLog(t *testing.T) {
	// t.Log("log")
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true) // 输出文件名
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("Info msg")
	logrus.Warning("Warning msg")
	logrus.Error("Error msg")
}

func TestFiledsLog(t *testing.T) {
	newLogger := logrus.WithFields(logrus.Fields{
		"ip":   "127.0.0.1",
		"port": 10013,
	}) // 新建一个携带固定内容的
	newLogger.Info("info test1 ")
	newLogger.Info("info test2 ")
}

func TestLogrusDesign(t *testing.T) {
	logrus.SetReportCaller(false)                // 配置可用
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 配置日志格式（logstash可用）
	logrus.Info("aaaaaa")
}

type AppHook struct {
	AppName string
}

func (ah *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (ah *AppHook) Fire(entry *logrus.Entry) error {
	// 这个方法在打印日志前调用
	entry.Data["app"] = ah.AppName
	return nil
}

// 设置 hook
func TestHook(t *testing.T) {
	h := &AppHook{AppName: "HelloWorldApp"}
	logrus.AddHook(h)
	logrus.Error("HOOK DEBUG")
}

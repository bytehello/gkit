package log

import (
	"github.com/gogap/logrus_mate"
	_ "github.com/gogap/logrus_mate/hooks/file"
	"github.com/sirupsen/logrus"
	"strings"
)

type GLogger struct {
	logger *logrus.Logger
	meta   *logrus_mate.LogrusMate
}

var (
	defaultLogger = New()
)

func New() *GLogger {
	return &GLogger{
		logger: logrus.StandardLogger(),
		meta:   newLogrusMeta(),
	}
}

func SetConfigString(configStr string, loggerName ...string) (*GLogger, error) {
	return setConfig(logrus_mate.ConfigString(configStr), loggerName...)
}

func SetConfigFile(fn string, loggerName ...string) (*GLogger, error) {
	return setConfig(logrus_mate.ConfigFile(fn), loggerName...)
}

func SetReportCaller(include bool) {
	defaultLogger.logger.SetReportCaller(include)
}

// Logger 参数不填写，则表示使用default（前提是传入的配置文件中有）
func Logger(loggerName ...string) *GLogger {
	return &GLogger{
		logger: defaultLogger.meta.Logger(loggerName...),
		meta:   newLogrusMeta(),
	}
}

func Debug(args ...interface{}) {
	defaultLogger.logger.Debug(args...)
}

func Info(args ...interface{}) {
	defaultLogger.logger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.logger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.logger.Error(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	defaultLogger.logger.Panic(args...)
}

func (gl *GLogger) Debug(args ...interface{}) {
	gl.logger.Debug(args...)
}

func (gl *GLogger) Info(args ...interface{}) {
	gl.logger.Info(args...)
}

func (gl *GLogger) Warn(args ...interface{}) {
	gl.logger.Warn(args...)
}

func (gl *GLogger) Error(args ...interface{}) {
	gl.logger.Error(args...)
}

func (gl *GLogger) Fatal(args ...interface{}) {
	gl.logger.Fatal(args...)
}

func (gl *GLogger) Panic(args ...interface{}) {
	gl.logger.Panic(args...)
}

func setConfig(opt logrus_mate.Option, loggerName ...string) (*GLogger, error) {
	name := "default"
	if len(loggerName) > 0 {
		name = strings.TrimSpace(loggerName[0])
		if len(name) == 0 {
			name = "default"
		}
	}

	meta := newLogrusMeta(opt)
	defaultLogger.meta = meta
	if err := meta.Hijack(defaultLogger.logger, name); err != nil {
		return nil, err
	}
	return defaultLogger, nil
}

func newLogrusMeta(opts ...logrus_mate.Option) *logrus_mate.LogrusMate {
	newMeta, _ := logrus_mate.NewLogrusMate(opts...)
	return newMeta
}

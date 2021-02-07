package log

import (
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus")


type GLogger struct {
	logger *logrus.Logger
	meta *logrus_mate.LogrusMate
}

var (
	defaultLogger = New()
)

func New() *GLogger {
	newMeta,_ := logrus_mate.NewLogrusMate()
	return &GLogger{
		logger: logrus.New(),
		meta: newMeta,
	}
}

func SetConfigString(configStr string) *GLogger {
	meta,_ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(configStr),
	)
	defaultLogger.meta = meta
	return defaultLogger
}

func (l *GLogger) Logger(loggerName ...string) *GLogger {
	l.logger = l.meta.Logger(loggerName ...)
	return l
}

func Debug(args ...interface{})  {
	defaultLogger.logger.Debug(args ...)
}

func Info(args ...interface{}) {
	defaultLogger.logger.Info(args ...)
}

func Warn(args ...interface{})  {
	defaultLogger.logger.Warn(args ...)
}

func Error(args ...interface{})  {
	defaultLogger.logger.Error(args ...)
}
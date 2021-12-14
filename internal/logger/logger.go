package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *logrus.Logger
}

var _ Interface = (*Logger)(nil)

func New(level string) *Logger {
	var l logrus.Level

	switch strings.ToLower(level) {
	case "error":
		l = logrus.ErrorLevel
	case "warn":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}
	logger := logrus.New()
	logger.SetLevel(l)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log("info", message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.log("warn", message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.msg("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(level string, message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf("%s - args: %v", message, args)
	}
	switch strings.ToLower(level) {
	case "debug":
		l.logger.Debug(message)
	case "info":
		l.logger.Info(message)
	case "warn":
		l.logger.Warn(message)
	case "error":
		l.logger.Error(message)
	case "fatal":
		l.logger.Fatal(message)
	default:
		l.logger.Info(message)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(level, fmt.Sprintf("%v of type %v", message, msg), args...)
	}
}

package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	log *logrus.Logger
}

func NewLogger(loglevel string) *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	switch loglevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
	return &Logger{
		log: log,
	}
}

func (l *Logger) Info(msg string, fields ...any) {
	l.log.Info(msg, fields)
}

func (l *Logger) Debug(msg string, fields ...any) {
	l.log.Debug(msg, fields)
}

func (l *Logger) Error(msg string, fields ...any) {
	l.log.Error(msg, fields)
}

func (l *Logger) Fatal(msg string, fields ...any) {
	l.log.Fatal(msg, fields)
}

func (l *Logger) Panic(msg string, fields ...any) {
	l.log.Panic(msg, fields)
}

func (l *Logger) Warn(msg string, fields ...any) {
	l.log.Warn(msg, fields)
}

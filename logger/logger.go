package logger

import (
	"context"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Init(logLevel string) {
	// SetFormatter
	log.SetFormatter(GetFormatter())

	// SetLevel
	switch strings.ToLower(logLevel) {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.SetOutput(os.Stdout)
}

func GetFormatter() log.Formatter {

	return &log.TextFormatter{
		FullTimestamp: true,
	}

}

func GetLogger() *log.Logger {
	return log.StandardLogger()
}

func Info(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value("X-Correlation-ID")
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		"Correlation-ID": id,
	}).Infof(format, values...)
}

func Warn(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value("X-Correlation-ID")
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		"Correlation-ID": id,
	}).Warnf(format, values...)
}

func Error(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value("X-Correlation-ID")
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		"Correlation-ID": id,
	}).Errorf(format, values...)
}

func Debug(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value("X-Correlation-ID")
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		"Correlation-ID": id,
	}).Debugf(format, values...)
}

func Fatal(ctx context.Context, format string, values ...interface{}) {
	var id string
	val := ctx.Value("X-Correlation-ID")
	if val != nil {
		id = val.(string)
	}
	log.WithFields(log.Fields{
		"Correlation-ID": id,
	}).Fatalf(format, values...)
}

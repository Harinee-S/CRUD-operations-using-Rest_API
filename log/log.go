package log

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	postgres = "POSTGRES"
	mongo    = "MONGO"
	redis    = "REDIS"
)
const (
	webRequest = "WEB_REQUEST"
	datastore  = "datastore"
)

type Logger struct {
	*logrus.Logger
}

var Log *Logger

type ErrorLogger struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		panic(ErrorLogger{Error: err})
	}
}

func Setup() {
	level, err := logrus.ParseLevel("INFO")
	panicIfError(err)

	logrusVars := &logrus.Logger{
		Out:       os.Stderr,
		Hooks:     make(logrus.LevelHooks),
		Formatter: &logrus.JSONFormatter{},
		Level:     level,
	}

	Log = &Logger{logrusVars}
}

func buildContext() logrus.Fields {
	return logrus.Fields{
		"application": map[string]string{
			"name":        "go-sample",
			"version":     "0.0.1",
			"environment": "development",
		},
	}
}

func (logger *Logger) httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.
		WithField("log_context", "web request").
		WithFields(buildContext()).
		WithFields(
			logrus.Fields{
				"context": webRequest,
				"request": map[string]interface{}{
					"path":          r.URL.Path,
					"requestMethod": r.Method,
					"host":          r.Host,
					"remoteAddress": r.RemoteAddr,
				},
			})
}

func (logger *Logger) HTTPErrorf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Errorf(format, args...)
}

func (logger *Logger) HTTPInfof(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Infof(format, args...)
}

func (logger *Logger) HTTPWarnf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Warnf(format, args...)
}

func (logger *Logger) persistenceLogEntry(dbType string) *logrus.Entry {
	var persistenceType string

	switch dbType {
	case postgres:
		persistenceType = "PostgreSQL"
		break
	case mongo:
		persistenceType = "MongoDB"
		break
	case redis:
		persistenceType = "REDIS"
		break
	default:
		persistenceType = "UNKNOWN"
		break
	}

	return logger.
		WithField("context", datastore).
		WithFields(buildContext()).
		WithFields(
			logrus.Fields{
				"type": persistenceType,
			})
}

func (logger *Logger) PostgresFatalf(format string, args ...interface{}) {
	logger.persistenceLogEntry(postgres).Fatalf(format, args...)
}

func (logger *Logger) PostgresErrorf(format string, args ...interface{}) {
	logger.persistenceLogEntry(postgres).Errorf(format, args...)
}

func (logger *Logger) PostgresInfof(format string, args ...interface{}) {
	logger.persistenceLogEntry(postgres).Infof(format, args...)
}

func (logger *Logger) PostgresWarnf(format string, args ...interface{}) {
	logger.persistenceLogEntry(postgres).Warnf(format, args...)
}

func (logger *Logger) MongoFatalf(format string, args ...interface{}) {
	logger.persistenceLogEntry(mongo).Fatalf(format, args...)
}

func (logger *Logger) MongoErrorf(format string, args ...interface{}) {
	logger.persistenceLogEntry(mongo).Errorf(format, args...)
}

func (logger *Logger) MongoInfof(format string, args ...interface{}) {
	logger.persistenceLogEntry(mongo).Infof(format, args...)
}

func (logger *Logger) MongoWarnf(format string, args ...interface{}) {
	logger.persistenceLogEntry(mongo).Warnf(format, args...)
}

func (logger *Logger) RedisFatalf(format string, args ...interface{}) {
	logger.persistenceLogEntry(redis).Fatalf(format, args...)
}

func (logger *Logger) RedisErrorf(format string, args ...interface{}) {
	logger.persistenceLogEntry(redis).Errorf(format, args...)
}

func (logger *Logger) RedisInfof(format string, args ...interface{}) {
	logger.persistenceLogEntry(redis).Infof(format, args...)
}

func (logger *Logger) RedisWarnf(format string, args ...interface{}) {
	logger.persistenceLogEntry(redis).Warnf(format, args...)
}

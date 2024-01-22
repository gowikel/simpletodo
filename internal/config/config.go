package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const (
	LOG_LEVEL_ENV = "LOG_LEVEL"
)

type LogLevel string

const (
	Debug    LogLevel = "DEBUG"
	Info     LogLevel = "INFO"
	Warn     LogLevel = "WARN"
	Error    LogLevel = "ERROR"
	Fatal    LogLevel = "FATAL"
	NoLevel  LogLevel = "NOLEVEL"
	Disabled LogLevel = "DISABLED"
	Trace    LogLevel = "TRACE"
)

type config struct {
	LogLevel zerolog.Level
}

var appConfig config

func init() {
	appConfig = load()
}

func load() config {
	result := config{}

	logLevel, err := getLogLevel()
	if err != nil {
		panic(err)
	}

	result.LogLevel = logLevel

	return result
}

func Config() config {
	return appConfig
}

func getLogLevel() (zerolog.Level, error) {
	level := strings.ToUpper(os.Getenv(LOG_LEVEL_ENV))

	switch level {
	case string(Debug):
		return zerolog.DebugLevel, nil
	case string(Info):
		return zerolog.InfoLevel, nil
	case string(Warn):
		return zerolog.WarnLevel, nil
	case string(Error):
		return zerolog.ErrorLevel, nil
	case string(Fatal):
		return zerolog.FatalLevel, nil
	case string(NoLevel):
		return zerolog.NoLevel, nil
	case string(Disabled):
		return zerolog.Disabled, nil
	case string(Trace):
		return zerolog.TraceLevel, nil
	}

	if len(level) > 0 {
		return zerolog.ErrorLevel, fmt.Errorf(
			"invalid log level: %q",
			level,
		)
	}

	return zerolog.ErrorLevel, nil
}

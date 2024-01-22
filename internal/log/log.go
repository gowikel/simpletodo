package log

import (
	"os"

	"github.com/gowikel/simpletodo/internal/config"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	config := config.Config()

	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	zerolog.SetGlobalLevel(config.LogLevel)
}

func GetLogger() zerolog.Logger {
	return logger
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Trace() *zerolog.Event {
	return logger.Trace()
}

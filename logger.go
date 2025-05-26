package logging

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	Standard *zerolog.Logger
}

var Log *Logger

func NewLogger() *Logger {
	logLevel := zerolog.InfoLevel
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{
		Standard: &logger,
	}
}

func init() {
	Log = NewLogger()
}

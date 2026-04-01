package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(level string) {
	parsedLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		parsedLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(parsedLevel)

	Log = zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout},
	).With().Timestamp().Caller().Logger()
}

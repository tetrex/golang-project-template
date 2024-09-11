package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(appEnv string) *zerolog.Logger {

	var output io.Writer
	switch appEnv {
	case "local":
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339} // FormatLevel: func(i interface{}) string {
	case "stage":
	default:
	}

	logger := zerolog.New(output).With().Timestamp().Logger()

	return &logger
}

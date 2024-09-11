package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func New(appEnv string) *zerolog.Logger {

	var output io.Writer
	switch appEnv {
	case "local":
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
			PartsExclude: []string{
				zerolog.TimestampFieldName,
			}}

	case "stage":
	default:
	}

	logger := zerolog.New(output).With().Timestamp().Logger()

	return &logger
}

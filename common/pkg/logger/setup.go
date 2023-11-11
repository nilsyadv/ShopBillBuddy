package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
)

// Init initializes the logger with the given log level and output.
// It should be called only once, in the main function.
func Init(level string, output string) (InterfaceLogger, error) {
	// Set the log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		log.Println("unknown log level defined. using default log level 'info'. expected log level value: debug, info, warn, error, fatal, panic")
	}

	logr := CustomeLogger{}

	if output == "stdout" && os.Getenv("ENV") != "production" {
		logr.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false}).With().Timestamp().Logger()
	} else {
		if output == "stdout" {
			logr.logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		} else if output == "stderr" {
			logr.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		} else {
			f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			logr.logger = zerolog.New(f).With().Timestamp().Logger()
		}
	}

	return &logr, nil
}

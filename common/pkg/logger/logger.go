package logger

import (
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/rs/zerolog"
)

type InterfaceLogger interface {
	Debug(msg string)
	Info(msg string)
	Infof(format string, v ...interface{})
	Warn(msg string)
	Error(msg string, err error)
	Fatal(msg string, err error)
	Panic(msg string, err error)
}

type CustomeLogger struct {
	logger zerolog.Logger
}

func InitLogger(conf config.InterfaceConfig) InterfaceLogger {
	return &CustomeLogger{}
}

// Debug logs a debug message.
func (custlg *CustomeLogger) Debug(msg string) {
	custlg.logger.Debug().Msg(msg)
}

// Info logs an info message.
func (custlg *CustomeLogger) Info(msg string) {
	custlg.logger.Info().Msg(msg)
}

// Info logs an info message.
func (custlg *CustomeLogger) Infof(format string, v ...interface{}) {
	custlg.logger.Info().Msgf(format, v...)
}

// Warn logs a warning message.
func (custlg *CustomeLogger) Warn(msg string) {
	custlg.logger.Warn().Msg(msg)
}

// Error logs an error message.
func (custlg *CustomeLogger) Error(msg string, err error) {
	custlg.logger.Err(err).Msg(msg)
}

// Fatal logs a fatal message and exits the program.
func (custlg *CustomeLogger) Fatal(msg string, err error) {
	custlg.logger.Fatal().Err(err).Msg(msg)
}

// Panic logs a panic message and panics.
func (custlg *CustomeLogger) Panic(msg string, err error) {
	custlg.logger.Panic().Err(err).Msg(msg)
}

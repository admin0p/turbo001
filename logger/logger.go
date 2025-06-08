package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func NewLogger() *slog.Logger {

	return slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	).With(
		slog.String("app", "realNow"),
		slog.Int("version", 1),
		slog.String("env", "development"),
	)
}

func init() {
	if log == nil {
		log = NewLogger()
	}
}

func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}

func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}
func Error(msg string, args ...any) {
	log.Error(msg, args...)
}

package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	log slog.Logger
}

func NewLogger() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}

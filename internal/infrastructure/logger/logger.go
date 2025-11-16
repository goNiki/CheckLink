package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	Log *slog.Logger
}

func NewLogger() *slog.Logger {
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		slog.Error("Failed to create logs directory", "error", err)
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	logFile, err := os.OpenFile(
		filepath.Join(logsDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log := slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}

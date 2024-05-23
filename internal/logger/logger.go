package logger

import (
	"log"
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envTest = "test"
	envProd = "prod"
)

func Init(env string) *slog.Logger {
	var l *slog.Logger

	switch env {
	case envDev:
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envTest:
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log.Fatal("Failed to initialize the logger, env: ", env)
	}

	return l
}

package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {

	const (
		envLocal = "local"
		envProd  = "prod"
	)

	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

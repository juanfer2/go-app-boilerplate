package utils

import (
	"time"

	"go.uber.org/zap"
)

func Logger(title string, err string) *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger.Warn(
		title,
		// Structured context as strongly typed Field values.
		zap.String("url", "url"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

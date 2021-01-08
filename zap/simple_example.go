package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync()

	zapLogger.Info("zap logger info",
		// Structured context as strongly typed Field values.
		zap.String("url", "asdfasdfasdfasdf"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
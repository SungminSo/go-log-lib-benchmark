package main

import (
	"fmt"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("check log level")

	log := zap.NewExample()
	defer log.Sync()

	log.Debug("debug level")
	log.Info("info level")
	log.Warn("warn level")
	log.Error("error level")
	// If the logger is in development mode, it then panics
	log.DPanic("dpanic level")
	// Calls os.Exit(1), even if logging at FatalLevel is disabled.
	log.Fatal("fatal level")
	// panics, even if logging at PanicLevel is disabled.
	log.Panic("panic level")
}
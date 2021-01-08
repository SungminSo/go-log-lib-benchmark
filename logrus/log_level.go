package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Will log anything that is info or above (warn, error, fatal, panic). Default.
	log.SetLevel(log.ErrorLevel)
	fmt.Println(log.GetLevel())
}

func main() {
	fmt.Println("check log level")

	// terminal에 딱히 안 찍힘
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")

	// terminal에 찍힘
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}

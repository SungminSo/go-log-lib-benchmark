package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(
		log.Fields{
			"key0": "value0",
			"key1": "value0",
			"key2": "value0",
			"key3": "value0",
			"key4": "value0",
			"key5": "value0",
		}).Info("simple log example")
}

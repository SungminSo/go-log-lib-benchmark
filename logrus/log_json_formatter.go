package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	startTime := time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		log.WithFields(
			logrus.Fields{
				"key0": "value0",
				"key1": "value0",
				"key2": "value0",
				"key3": "value0",
				"key4": "value0",
				"key5": "value0",
			}).Info("INFO level")

		log.WithFields(
			logrus.Fields{
				"key0": "value0",
				"key1": "value0",
				"key2": "value0",
				"key3": "value0",
				"key4": "value0",
				"key5": "value0",
			}).Warn("WARN level")

		log.WithFields(
			logrus.Fields{
				"key0": "value0",
				"key1": "value0",
				"key2": "value0",
				"key3": "value0",
				"key4": "value0",
				"key5": "value0",
			}).Error("ERROR level")
	}


	// 총 2000 돌렸을 때 걸리는 시간: 평균 228674000 정도
	// 총 20000 돌렸을 때 걸리는 시간: 평균 2197045000 정도
	// 필드 중에 "level"이 있어서 구분은 가능하지만 눈에 띄지는 않음
	// 필드 중 "msg", "time"이 있음
	fmt.Printf("logging finished: %v", time.Now().UnixNano() - startTime)
}
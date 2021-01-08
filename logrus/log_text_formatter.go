package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{})

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

	// 총 2000 돌렸을 때 걸리는 시간: 평균 977067000 정도
	// 총 20000 돌렸을 때 걸리는 시간: 평균 9649492000 정도
	// level이 색상때문에 눈에 잘 띔
	// 필드 중 time을 따로 넣어줘야 함
	fmt.Printf("logging finished: %v", time.Now().UnixNano() - startTime)
}
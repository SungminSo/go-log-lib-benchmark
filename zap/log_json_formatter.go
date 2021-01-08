package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	startTime := time.Now().UnixNano()
	for i := 0; i < 2000; i++ {
		log.Info("INFO level",
			zap.String("key0", "value0"),
			zap.String("key1", "value0"),
			zap.Int("key2", 22),
			zap.Int("key3", 333),
			zap.Duration("key4", time.Second),
			zap.Duration("key5", time.Millisecond),
		)

		log.Warn("WARN level",
			zap.String("key0", "value0"),
			zap.String("key1", "value0"),
			zap.Int("key2", 22),
			zap.Int("key3", 333),
			zap.Duration("key4", time.Second),
			zap.Duration("key5", time.Millisecond),
		)

		log.Error("Error level",
			zap.String("key0", "value0"),
			zap.String("key1", "value0"),
			zap.Int("key2", 22),
			zap.Int("key3", 333),
			zap.Duration("key4", time.Second),
			zap.Duration("key5", time.Millisecond),
		)
	}

	// 총 2000 돌렸을 때 걸리는 시간: 평균 27661000 정도
	// 총 20000 돌렸을 때 걸리는 시간: 평균 58721000 정도
	// Error 일때는 "stacktrace" 필드가 추가로 있음
	fmt.Printf("logging finished: %v", time.Now().UnixNano() - startTime)
}

package utils

import (
	"math/rand"
	"time"
)

func JobShort() int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return 1
}

func JobLong() int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return 1
}

func GenerateId() string {
	currentTime := time.Now()
	return currentTime.Format(time.RFC3339Nano)
}

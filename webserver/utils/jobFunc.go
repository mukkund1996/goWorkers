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
	return 5
}

func GenerateId(prefix string) string {
	currentTime := time.Now()
	return prefix + "_" + currentTime.Format(time.RFC3339Nano)
}

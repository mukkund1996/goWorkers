package utils

import (
	"math/rand"
	"time"
	"webserver/models"
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

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func CollectResults(receiver <-chan models.ResultSpec, results map[string][]int, jobs *[]models.JobSpec) {
	for r := range receiver {
		*jobs = filter(*jobs, func(j models.JobSpec) bool { return j.Id == r.Id })
		result := results[r.Id]
		result = append(result, r.Result)
		results[r.Id] = result
	}
}

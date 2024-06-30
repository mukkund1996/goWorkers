package controllers

import (
	"strconv"
	"webserver/models"
	"webserver/utils"

	"github.com/gin-gonic/gin"
)

func ServerHealthCheck(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}

func RunShortJobs(s chan models.JobSpec, r chan int) func(c *gin.Context) {
	return func(c *gin.Context) {
		results := []int{}
		numWorkers, err := strconv.Atoi(c.Param("workerCount"))
		if err != nil {
			return
		}

		for i := 0; i < numWorkers; i++ {
			jobSpec := models.JobSpec{Id: utils.GenerateId(), Operation: utils.JobShort}
			s <- jobSpec
		}
		for i := 0; i < numWorkers; i++ {
			results = append(results, <-r)
		}
		c.IndentedJSON(200, gin.H{
			"status": "Completed",
			"result": results,
		})
	}
}

func RunLongJobs(s chan models.JobSpec, r chan int) func(c *gin.Context) {
	return func(c *gin.Context) {
		results := []int{}
		numWorkers, err := strconv.Atoi(c.Param("workerCount"))
		if err != nil {
			return
		}

		for i := 0; i < numWorkers; i++ {
			jobSpec := models.JobSpec{Id: utils.GenerateId(), Operation: utils.JobLong}
			s <- jobSpec
		}
		for i := 0; i < numWorkers; i++ {
			results = append(results, <-r)
		}
		c.IndentedJSON(200, gin.H{
			"status": "Completed",
			"result": results,
		})
	}
}

func CheckWorkerStatus(workers []models.Worker) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"workers": workers,
		})
	}
}

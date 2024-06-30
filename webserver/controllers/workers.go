package controllers

import (
	"encoding/json"
	"strconv"
	"time"
	"webserver/config"
	"webserver/models"
	"webserver/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

func WorkerStatusSocket(workers *[]models.Worker) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			message := []byte("")
			for w := range *workers {
				body, _ := json.Marshal(w)
				message = append(message, []byte(body)...)
			}
			conn.WriteMessage(websocket.TextMessage, message)
			time.Sleep(config.SocketPollingInterval)
		}
	}
}

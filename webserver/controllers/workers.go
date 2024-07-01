package controllers

import (
	"encoding/json"
	"net/http"
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
	CheckOrigin:     func(*http.Request) bool { return true },
}

func ServerHealthCheck(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}

func RunShortJobs(s chan models.JobSpec, jobs *[]models.JobSpec) func(c *gin.Context) {
	return func(c *gin.Context) {
		numWorkers, err := strconv.Atoi(c.Param("workerCount"))
		if err != nil {
			return
		}
		reqId := utils.GenerateId("short")
		for i := 0; i < numWorkers; i++ {
			jobSpec := models.JobSpec{Id: reqId, Operation: utils.JobShort}
			*jobs = append(*jobs, jobSpec)
			s <- jobSpec
		}
		c.IndentedJSON(200, gin.H{
			"status": "Started",
		})
	}
}

func RunLongJobs(s chan models.JobSpec, jobs *[]models.JobSpec) func(c *gin.Context) {
	return func(c *gin.Context) {
		numWorkers, err := strconv.Atoi(c.Param("workerCount"))
		if err != nil {
			return
		}

		for i := 0; i < numWorkers; i++ {
			jobSpec := models.JobSpec{Id: utils.GenerateId("long"), Operation: utils.JobLong}
			*jobs = append(*jobs, jobSpec)
			s <- jobSpec
		}
		c.IndentedJSON(200, gin.H{
			"status": "Completed",
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
			body, _ := json.Marshal(workers)
			conn.WriteMessage(websocket.TextMessage, body)
			time.Sleep(config.SocketPollingInterval)
		}
	}
}

func JobResultSocket(results map[string][]int) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			body, _ := json.Marshal(results)
			conn.WriteMessage(websocket.TextMessage, body)
			time.Sleep(config.SocketPollingInterval)
		}
	}
}

func JobQueueSocket(jobs []models.JobSpec) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			ids := []string{}
			for _, j := range jobs {
				ids = append(ids, j.Id)
			}
			body, _ := json.Marshal(ids)
			conn.WriteMessage(websocket.TextMessage, body)
			time.Sleep(config.SocketPollingInterval)
		}
	}
}

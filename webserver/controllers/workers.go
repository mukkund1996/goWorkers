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

func RunShortJobs(s chan models.JobSpec, jobs *[]string) func(c *gin.Context) {
	return func(c *gin.Context) {
		numWorkers, err := strconv.Atoi(c.Param("workerCount"))
		if err != nil {
			return
		}
		reqId := utils.GenerateId("short")
		c.IndentedJSON(200, gin.H{
			"status": "Submitted",
			"job":    reqId,
		})
		for i := 0; i < numWorkers; i++ {
			*jobs = append(*jobs, reqId)
		}
		for i := 0; i < numWorkers; i++ {
			jobSpec := models.JobSpec{Id: reqId, Operation: utils.JobShort}
			s <- jobSpec
			*jobs = (*jobs)[:len(*jobs)-1]
		}
	}
}

func RunLongJobs(s chan models.JobSpec, jobs *[]string) func(c *gin.Context) {
	return func(c *gin.Context) {
		numWorkers, err := strconv.Atoi(c.Param("workerCount"))
		if err != nil {
			return
		}

		reqId := utils.GenerateId("long")
		c.IndentedJSON(200, gin.H{
			"status": "Submitted",
			"job":    reqId,
		})
		for i := 0; i < numWorkers; i++ {
			*jobs = append(*jobs, reqId)
		}
		for i := 0; i < numWorkers; i++ {
			jobSpec := models.JobSpec{Id: reqId, Operation: utils.JobLong}
			s <- jobSpec
			*jobs = (*jobs)[:len(*jobs)-1]
		}
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

func JobQueueSocket(jobs *[]string) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			body, _ := json.Marshal(jobs)
			conn.WriteMessage(websocket.TextMessage, body)
			time.Sleep(config.SocketPollingInterval)
		}
	}
}

package routes

import (
	"webserver/controllers"
	"webserver/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter(s chan models.JobSpec, r chan models.ResultSpec, workers []models.Worker, jobs *[]models.JobSpec, results map[string][]int) *gin.Engine {
	router := gin.Default()

	// Websocket handlers
	router.GET("/queue", controllers.JobQueueSocket(*jobs))
	router.GET("/status", controllers.CheckWorkerStatus(workers))
	router.GET("/workerStatus", controllers.WorkerStatusSocket(&workers))

	// HTTP request handlers
	router.GET("/ping", controllers.ServerHealthCheck)
	router.GET("/results", controllers.JobResultSocket(results))
	router.POST("/short/:workerCount", controllers.RunShortJobs(s, jobs))
	router.POST("/long/:workerCount", controllers.RunLongJobs(s, jobs))

	return router
}

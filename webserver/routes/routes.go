package routes

import (
	"webserver/controllers"
	"webserver/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter(s chan models.JobSpec, r chan int, workers []models.Worker) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.ServerHealthCheck)

	router.GET("/status", controllers.CheckWorkerStatus(workers))

	router.GET("/ws", controllers.WorkerStatusSocket(&workers))

	router.POST("/short/:workerCount", controllers.RunShortJobs(s, r))

	router.POST("/long/:workerCount", controllers.RunLongJobs(s, r))

	return router
}

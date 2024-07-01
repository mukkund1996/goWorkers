package routes

import (
	"webserver/controllers"
	"webserver/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter(s chan models.JobSpec, r chan int, workers []models.Worker, jobs *[]models.JobSpec) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.ServerHealthCheck)

	router.GET("/status", controllers.CheckWorkerStatus(workers))

	router.GET("/ws", controllers.WorkerStatusSocket(&workers))

	router.POST("/short/:workerCount", controllers.RunShortJobs(s, r, jobs))

	router.POST("/long/:workerCount", controllers.RunLongJobs(s, r, jobs))

	return router
}

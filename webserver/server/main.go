package main

import (
	"fmt"
	"webserver/config"
	"webserver/models"
	"webserver/routes"
)

func main() {
	sender := make(chan models.JobSpec, config.MaxQueueLength)
	receiver := make(chan int, config.MaxQueueLength)

	// Creating and initializing workers
	workers := make([]models.Worker, config.DefaultNumWorkers)

	for i := range workers {
		workers[i] = models.CreateWorker(i, sender, receiver)
	}

	// Starting the workers
	for i := range workers {
		go workers[i].StartListening()
	}

	router := routes.SetupRouter(sender, receiver, workers)
	router.Run(fmt.Sprintf(":%d", config.AppPort))
}

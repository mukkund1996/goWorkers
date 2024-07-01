package main

import (
	"fmt"
	"webserver/config"
	"webserver/models"
	"webserver/routes"
	"webserver/utils"
)

func main() {
	sender := make(chan models.JobSpec, config.MaxQueueLength)
	receiver := make(chan models.ResultSpec, config.MaxQueueLength)
	jobs := []models.JobSpec{}
	results := make(map[string][]int)

	// Creating and initializing workers
	workers := make([]models.Worker, config.DefaultNumWorkers)

	for i := range workers {
		workers[i] = models.CreateWorker(i, sender, receiver)
	}

	// Starting the workers
	for i := range workers {
		go workers[i].StartListening()
	}

	// Consolidate the results
	go utils.CollectResults(receiver, results, &jobs)

	router := routes.SetupRouter(sender, receiver, workers, &jobs, results)
	router.Run(fmt.Sprintf(":%d", config.AppPort))
}

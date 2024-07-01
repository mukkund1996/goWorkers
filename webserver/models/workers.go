package models

import (
	"fmt"
)

type JobSpec struct {
	Id        string
	Operation func() int
}

type ResultSpec struct {
	Id     string
	Result int
}

type Worker struct {
	Id    int
	Busy  *bool
	JobId *string
	r     <-chan JobSpec
	s     chan<- ResultSpec
}

func (w Worker) StartListening() {
	fmt.Printf("Worker %d started listening to jobs...\n", w.Id)
	for j := range w.r {
		fmt.Printf("Worker %d: Received job %s\n", w.Id, j.Id)
		*(w.Busy) = true
		*(w.JobId) = j.Id
		w.s <- ResultSpec{Id: j.Id, Result: j.Operation()}
		fmt.Printf("Worker %d: Completed job\n", w.Id)
		*(w.Busy) = false
		*(w.JobId) = ""
	}
}

func CreateWorker(i int, r <-chan JobSpec, s chan<- ResultSpec) Worker {
	isWorking := false
	job := ""
	return Worker{Id: i, Busy: &isWorking, JobId: &job, r: r, s: s}
}

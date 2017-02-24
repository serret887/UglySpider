package workerPool

import (
	"fmt"
	"strconv"
)

//Worker methods to manage a pile of workers
type Worker interface {
	Start()
	Close() error
}

//Job can mock any method that need to be executed
type Job interface {
	Execute()
}

var count int

type worker struct {
	WorkerName     string
	WorkerPool     chan chan *Job
	RequestChannel chan *Job
	quit           chan chan error
}

// NewWorker return a new minion that will execute any job for you
func NewWorker(workerPool chan chan *Job) Worker {
	count++
	return &worker{
		WorkerName:     "baby_minion" + strconv.Itoa(count),
		WorkerPool:     workerPool,
		RequestChannel: make(chan *Job),
		quit:           make(chan chan error),
	}
}

// Start method starts the run loop for the worker, listening for a
// quit chanel in case we stop the request
func (w *worker) Start() {
	go func() {
		var err error
		for {
			//Register this worker in to the queue
			w.WorkerPool <- w.RequestChannel
			select {
			case job := <-w.RequestChannel:
				// in this case we receive a job so we need to execute it
				fmt.Println("job executing", job)
				// we need to send the response from job to his own channel
			case q := <-w.quit:
				//this is the stop signal
				fmt.Println("Stopping the worker")
				q <- err
				return
			}
		}
	}()
}

func (w *worker) Close() error {
	errc := make(chan error)
	w.quit <- errc
	return <-errc
}

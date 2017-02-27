package workerPool

import "fmt"

//Worker methods to manage a pile of workers
type Worker interface {
	Start()
	Close() error
}

//Job can mock any method that need to be executed
type Job interface {
	Execute() error
	fmt.Stringer
	GetResult() (interface{}, error)
}

type worker struct {
	WorkerName     string
	WorkerPool     chan chan *Job
	RequestChannel chan *Job
	ResponseChan   chan error
	quit           chan chan error
}

// NewWorker return a new minion that will execute any job for you
func NewWorker(workerPool chan chan *Job, RespChan chan error, name string) Worker {
	return &worker{
		WorkerName:     name,
		WorkerPool:     workerPool,
		RequestChannel: make(chan *Job),
		ResponseChan:   RespChan,
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
			fmt.Println("Worker registered")
			select {
			case job := <-w.RequestChannel:
				// in this case we receive a job so we need to execute it
				fmt.Println("the job wil be execute")

				err := (*job).Execute()
				// we need to send the response from job to his own channel
				fmt.Println("job executed")
				w.ResponseChan <- err
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
	fmt.Println(w.WorkerName, "stopped")
	return <-errc
}

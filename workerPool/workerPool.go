package workerPool

var (
	maxWorkers = 2 // os.Getenv("MAX_WORKERS")
	maxQueue   = 2 //os.Getenv("MAX_QUEUE")
)

// Dispatcher is in charge of create a pool of jobs
type Dispatcher interface {
	Dispatch()
}
type dispatcher struct {
	// pool of worker that is trigger and registered with worker
	WorkPool chan chan *Job
	//JObInput is for send a job to the pool
	JobInput chan *Job
	// Job Output is also need it here

}

// NewJobDispatcher return a new JOB dispatcher
func NewJobDispatcher() Dispatcher {
	return &dispatcher{
		WorkPool: make(chan chan *Job, maxWorkers),
		JobInput: make(chan *Job),
	}
}

func (d *dispatcher) Dispatch() {
	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(d.WorkPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobInput:
			go func(job *Job) {
				// somebody throw a job to the work pool
				// pick one worker from the queue
				jobChan := <-d.WorkPool
				// dispatch the job to the worker channel
				jobChan <- job
			}(job)
		}
	}
}

package workerPool

import (
	"fmt"
	"strconv"
	"sync"
)

// Dispatcher is in charge of create a pool of jobs
type Dispatcher interface {
	Dispatch(workerAmount int)
	GetJobInput() chan *Job
	GetResponse() chan error
	Close() error
	GetAmountOfWorkers() int
}

type dispatcher struct {
	PoolName   string
	workerList []*Worker
	// pool of worker that is trigger and registered with worker
	WorkPool chan chan *Job
	//JObInput is for send a job to the pool
	JobInput chan *Job
	// Job Output is also need it here
	ResponseChan chan error
	closing      chan chan error
}

// NewJobDispatcher return a new JOB dispatcher
func NewJobDispatcher(name string) Dispatcher {
	return &dispatcher{
		PoolName: name,
		JobInput: make(chan *Job),
		closing:  make(chan chan error),
	}
}

func (d *dispatcher) Dispatch(amountWorkers int) {
	d.WorkPool = make(chan chan *Job, amountWorkers)
	d.workerList = make([]*Worker, amountWorkers)
	d.ResponseChan = make(chan error, amountWorkers)
	fmt.Print(d.WorkPool)
	for i := 0; i < amountWorkers; i++ {
		workerName := d.PoolName + "_minion_" + strconv.Itoa(i+1)
		worker := NewWorker(d.WorkPool, d.ResponseChan, workerName)
		worker.Start()
		d.workerList[i] = &worker
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
				fmt.Print(*job)
				fmt.Print(jobChan)
				fmt.Println("The job is send to executed")
				//	runtime.Breakpoint()
				fmt.Println(*job)
				jobChan <- job
				fmt.Println("the job was wxecuted")
				// worker execute the job and then we get the response
			}(job)

		case q := <-d.closing:

			// iterate between the worker calling close
			// blocking the jobInput so the select is blocked
			var err error
			// go func() {
			// 	d.JobInput = nil
			// }()

			fmt.Println("closing primero")
			var wg sync.WaitGroup
			wg.Add(d.GetAmountOfWorkers())
			fmt.Println(d.GetAmountOfWorkers())
			for i, w := range d.workerList {
				go func(w *Worker, wg *sync.WaitGroup, i int) {
					defer wg.Done()
					fmt.Println("closing", i)
					// can put a list of errors for the workers
					(*w).Close()
				}(w, &wg, i)
			}
			fmt.Println("waiting for closing")
			wg.Wait()
			fmt.Println("no more waiting for closing")
			// close the worker Pool
			fmt.Print("WORK POOL CLOSE")
			q <- err
			return

		}

	}
}

func (d *dispatcher) GetJobInput() chan *Job {
	return d.JobInput
}

func (d *dispatcher) GetResponse() chan error {
	return d.ResponseChan
}

func (d *dispatcher) Close() error {
	errc := make(chan error)
	d.closing <- errc
	return <-errc
}

func (d *dispatcher) GetAmountOfWorkers() int {
	return len(d.workerList)
}

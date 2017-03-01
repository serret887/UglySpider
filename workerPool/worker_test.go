package workerPool_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/UglySpider/workerPool"
)

type DummyJob struct {
	Count int
}

func (dj *DummyJob) Execute() error {
	dj.Count = 5
	return nil
}

func (dj *DummyJob) String() string {
	return fmt.Sprint("DUMMY JOB")
}

func (dj *DummyJob) GetResult() (interface{}, error) {
	return dj.Count, nil
}

func (dj *DummyJob) Close() error {
	return nil
}

var _ = Describe("Worker", func() {
	Context("Worker Test", func() {
		It("Worker should execute any type that implement Excute interface", func() {
			wp := workerPool.NewJobDispatcher("test1")
			wp.Dispatch(1)

			var dj workerPool.Job
			dj = &DummyJob{}

			// assing job to the queue
			wp.GetJobInput() <- &dj

			resp := <-wp.GetResponse()
			Expect(resp).To(BeNil())
			djStruct, ok := dj.(*DummyJob)
			Expect(ok).To(BeTrue())
			Expect(djStruct.GetResult()).To(Equal(5), "expecting to execute the job")
			Expect(dj.String()).To(Equal("DUMMY JOB"), "expectiong to receive the same job that i send")
			err := wp.Close()
			Expect(err).To(BeNil())
		})

		It("Should dispatch the amount of worker required", func() {
			wp := workerPool.NewJobDispatcher("test")
			wp.Dispatch(3)
			Expect(wp.GetAmountOfWorkers()).To(Equal(3), "the amount of workers should be 3")
			err := wp.Close()
			Expect(err).To(BeNil())

		})

	})

})

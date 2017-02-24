package workerPool_test

import (
	"github.com/serret887/UglySpider/workerPool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DummyJob struct {
	Count int
}

func (dj *DummyJob) Execute() {
	dj.Count++
}

var _ = Describe("Worker", func() {

	It("Worker should execute any type that implement Excute interface", func() {
		wp := make(chan chan *workerPool.Job)
		w := workerPool.NewWorker(wp)
		w.Start()
		dj := &DummyJob{Count: 1}
		w <- dj

	})
})

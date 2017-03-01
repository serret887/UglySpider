package UglySpider

import (
	"github.com/serret887/UglySpider/workerPool"
	"github.com/serret887/ogle/matcher"
)

// Scrapper manage all the workers and set the rules
// for scrappe
type Scrapper struct {
	Domain     string
	link, data matcher.Matcher
	wp         workerPool.Dispatcher
	job        *workerPool.Job
}

// NewScrapper is the constructor of a new scrapper it
// create a worker pool with the domain name and manage
// all the workers for jobs
func NewScrapper(domain string, workers int, task *workerPool.Job) *Scrapper {
	pool := workerPool.NewJobDispatcher("Fetch_" + domain)
	pool.Dispatch(workers)
	return &Scrapper{
		Domain: domain,
		wp:     pool,
		job:    task,
	}
}

//
func (s *Scrapper) Fetch() interface{} {
	return nil
}

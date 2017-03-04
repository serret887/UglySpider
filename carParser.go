package UglySpider

import (
	"github.com/serret887/UglySpider/workerPool"
)

// TODO Create a version rule by a file of rules
// implemented in xml with keywords that i need to define

// JobParser is a way to comunicate and do parsing jobs
type JobParser interface {
	workerPool.Job
	GetResult() //identify what will be the response
}
type jobParser struct {
}

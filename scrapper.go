package UglySpider

import (
	"github.com/serret887/ogle/matcher"
)

// Scrapper manage all the workers and set the rules
// for scrappe
type Scrapper struct {
	Domain     string
	link, data matcher.Matcher
}

func (s *Scrapper) Fetch(url string, r bool) {
	if r == true {
		go s.loop()
		return
	}

}

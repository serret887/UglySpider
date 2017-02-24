package coach

import (
	"github.com/serret887/UglySpider/pitcher"
	"github.com/serret887/ogle"
)

//Coach is the pitchers managers is given to him a task
//ex: extract all the links from this website
// his responsability is create a request for the pitchers to do
// and parse this request in reference to the matchers passed to him in his creation
// after that he will notify his response and will wait for more duties to perform.
type Coach struct {
	pitchers pitcher.Pitcher
}

// NewCoach make a new coach with some team players
func NewCoach(pAddress string) (*Coach, error) {
	coach := &Coach{}
	//for i := 0; i < n; i++ {
	p, err := pitcher.NewPitcher(pAddress)
	if err != nil {
		return nil, err
	}
	coach.pitchers = *p
	//}
	coach.urlChan = make(chan string, n)
	return coach, nil
}

func (c *Coach) Process(urls <-chan string) *ogle.Ogle {
for _, url := range urls{
	select{

	}
}
	}

}

func (c *Coach) Close() error {
	errc := make(chan error)
	c.closing <- errc
	return <-errc
}

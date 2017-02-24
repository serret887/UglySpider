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
	pitchers []pitcher.Pitcher
}

// NewCoach make a new coach with some team players
func NewCoach(n int, pAddress string) (*Coach, error) {
	coach := &Coach{}
	for i := 0; i < n; i++ {
		p, err := pitcher.NewPitcher(pAddress)
		if err != nil {
			return nil, err
		}
		coach.pitchers = append(coach.pitchers, *p)
	}
	return coach, nil
}

func (c *Coach) Process(ogleChan chan []*ogle.Ogle, urls ...string) {

}

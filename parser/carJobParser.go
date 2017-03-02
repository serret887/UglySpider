package parser

import (
	"fmt"

	"io"

	"github.com/serret887/UglySpider/workerPool"
	"github.com/serret887/ogle"
	"github.com/serret887/ogle/matcher"
	"golang.org/x/net/html/atom"
)

// CarJobParser is specifyc job for the cars
// this is the first implementation and it will be repalce for
// one more flexible
type CarJobParser interface {
	workerPool.Job
	GetResult()
	SetDataInput(io.Reader) error
}

type carJobParser struct {
	Price, Place, Details []matcher.Matcher
	ogleData              *ogle.Ogle
}

// NewCarJobParser constructor for this stuff
func NewCarJobParser() (CarJobParser, error) {

	parentCommon := matcher.WithParent(
		matcher.WithTag(atom.Span),
		matcher.WithClass("result-meta"))

	resultPriceMatcher := []matcher.Matcher{
		matcher.WithTag(atom.Span),
		matcher.WithClass("result-price"),
	}
	resultPriceMatcher = append(resultPriceMatcher, parentCommon)
	// place matcher
	placeMatcher := []matcher.Matcher{
		matcher.WithTag(atom.Span),
		matcher.WithClass("result-hood"),
	}
	placeMatcher = append(placeMatcher, parentCommon)
	// details page this is for the user to review
	detailPageMatcher := []matcher.Matcher{
		matcher.WithTag(atom.A),
		matcher.WithClass("result-title hdrlnk"),
		matcher.WithParent(matcher.WithTag(atom.Span))}
	return &carJobParser{
		Price:   resultPriceMatcher,
		Place:   placeMatcher,
		Details: detailPageMatcher,
	}, nil
}

// Execute make the car list from the ogle and the matchers
func (cj *carJobParser) Execute() error {
	price, err := cj.ogleData.Find(cj.Price...)
	if err != nil {
		return err
	}
	// detail, err := cj.ogleData.Find(cj.Details...)
	// if err != nil {
	// 	return err
	// }
	// place, err := cj.ogleData.Find(cj.Place...)
	// if err != nil {
	// 	return err
	// }
	fmt.Print("amount of price in the page", len(price))
	for _, pr := range price {
		fmt.Print(pr.FirstChild)
	}

	return nil
}

func (cj *carJobParser) SetDataInput(data io.Reader) error {
	var err error
	cj.ogleData, err = ogle.New(data)
	return err
}

func (cj *carJobParser) Close() error {
	//cj.ogleData = nil
	// wait to see what happend
	return nil
}
func (cj *carJobParser) String() string {
	return fmt.Sprintf("parsing for node with %s and node with %s and node with %s", cj.Place, cj.Price, cj.Details)
}

func (cj *carJobParser) GetResult() {

}

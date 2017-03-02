package parser_test

import (
	"github.com/serret887/UglySpider/parser"

	"bytes"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CarJobParser", func() {
	// loading content from a file to a io.Reader
	f, err := os.Open("../Resources/craig.html")
	if err != nil {
		panic("Maybe the resource is missing ")
	}

	Context("Sending data to the parser get parser succesfully", func() {
		It("SetDataInput does not return any error with valid data", func() {
			p, err := parser.NewCarJobParser()
			Expect(err).To(BeNil(), "no error is expected")
			var buf bytes.Buffer
			err = p.SetDataInput(&buf)
			Expect(err).To(BeNil(), "no error is expected")
		})

		It("return the right description when the String method is called", func() {
			p, err := parser.NewCarJobParser()
			Expect(err).To(BeNil(), "no error is expected")
			Expect(p.String()).To(BeEquivalentTo("parsing for node with [tag span class result-hood ] and node with [tag span class result-price ] and node with [tag a class result-title hdrlnk ]"))
		})

		It("Return all the car data from the query page", func() {
			p, err := parser.NewCarJobParser()
			Expect(err).To(BeNil(), "no error is expected")
			err = p.SetDataInput(f)
			Expect(err).To(BeNil(), "no error is expected for io.reader")
			err = p.Execute()
			Expect(err).To(BeNil(), "no error is expected for whe executing the job")

		})

	})
})

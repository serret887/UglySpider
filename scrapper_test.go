package carScrapper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scrapper", func() {

	Context("make html request", func() {
		It("Requesting https://detroit.craigslist.org/search/cta?s=100 get all the links for cars", func() {
			//Scrapper{Domain: `https://detroit.craigslist.org/search/cta?s=100`}
			//	links := sc.getLinks()
			Expect(1).To(Equal(100), "every page in craighlist have 100 cars")
		})
	})

})

package pitcher_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/UglySpider/pitcher"
)

var _ = Describe("Pitcher", func() {
	Context("making HTTP request", func() {
		var fakeServer *httptest.Server
		BeforeEach(func() {
			fakeServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cookie, err := r.Cookie("times")
				if err != nil {
					cookie = &http.Cookie{Name: "times"}
				}
				fmt.Print("Request executed")
				cookie.Value += "1"
				http.SetCookie(w, cookie)
				fmt.Fprintln(w, "Hello, client")
			}))
		})

		It("Will get the correct server response", func() {

			p, err := pitcher.NewPitcher(fakeServer.URL)
			Expect(err).To(BeNil(), "Should be able to create a Pitcher")
			res, err := p.Get(fakeServer.URL)
			Expect(err).To(BeNil(), "Should make a get request without problems")
			defer res.Body.Close()
			Expect(res).NotTo(BeEquivalentTo("Hello Client"), "the response should keep the cookies")
		})
		It("Will save the cookies passed by the server", func() {

			p, err := pitcher.NewPitcher(fakeServer.URL)
			Expect(err).To(BeNil(), "Should be able to create a Pitcher")
			res, err := p.Get(fakeServer.URL)
			Expect(err).To(BeNil(), "Should make a get request without problems")
			defer res.Body.Close()
			cookies := res.Cookies()
			Expect(cookies[0].Name).To(Equal("times"), "the name of the cookie should be times")
			Expect(cookies[0].Value).To(Equal("1"), "the name of the cookie should be times")

		})
		It("Will save the cookies passed by the server for all the request", func() {

			p, err := pitcher.NewPitcher(fakeServer.URL)
			Expect(err).To(BeNil(), "Should be able to create a Pitcher")
			res, err := p.Get(fakeServer.URL)
			res, err = p.Get(fakeServer.URL)
			res, err = p.Get(fakeServer.URL)
			Expect(err).To(BeNil(), "Should make a get request without problems")
			defer res.Body.Close()
			cookies := res.Cookies()
			Expect(cookies[0].Name).To(Equal("times"), "the name of the cookie should be times")
			Expect(cookies[0].Value).To(Equal("111"), "the name of the cookie should be times")

		})

		AfterEach(func() {
			fakeServer.Close()
		})
	})
})

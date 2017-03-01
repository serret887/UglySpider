package UglySpider_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/UglySpider"
	"github.com/serret887/UglySpider/pitcher"
)

var _ = Describe("Jobrequest", func() {
	Context("Jobrequest can live in his own enviroment", func() {

		It("When a job request is create it close without problems", func() {
			job, err := UglySpider.NewJobRequest("")

			Expect(err).To(BeNil(), "no errors creating the job")
			Expect(job.Close()).To(BeNil(), "no error when closing the request")
		})
	})

	Context("making HTTP request", func() {
		var fakeServer *httptest.Server
		var request *http.Request

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

		AfterEach(func() {
			fakeServer.Close()
		})

		JustBeforeEach(func() {
			var err error
			request, err = http.NewRequest("GET", fakeServer.URL, nil)
			Expect(err).To(BeNil(), "Should be able to create a Pitcher")
		})

		FIt("Will get the correct response from a server", func() {
			rj, err := UglySpider.NewJobRequest(fakeServer.URL)
			Expect(err).To(BeNil(), "Should be able to create a Request")
			rj.SetRequest(request)
			//			runtime.Breakpoint()
			err = rj.Execute()
			Expect(err).To(BeNil(), " make a get request without problems")
			response, err := rj.GetResult()
			Expect(err).To(BeNil(), "Should be able to create a Request")
			defer response.Body.Close()
			resp, err := ioutil.ReadAll(response.Body)
			Expect(err).To(BeNil(), "should be no error in the response")
			runtime.Breakpoint()
			Expect(resp).To(BeEquivalentTo([]byte("Hello, client")), "the response should keep the cookies")

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

	})
})

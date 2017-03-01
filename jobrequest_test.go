package UglySpider_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/UglySpider"
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
				fmt.Println("Request executed")
				fmt.Println(cookie)
				cookie.Value += "1"
				fmt.Println(cookie)
				http.SetCookie(w, cookie)
				fmt.Fprint(w, "Hello, client")
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

		It("Will get the correct response from a server", func() {
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

			Expect(string(resp)).To(BeEquivalentTo("Hello, client"), "the response should should be in the body")

		})

		It("Will save the cookies passed by the server", func() {

			rj, err := UglySpider.NewJobRequest(fakeServer.URL)
			Expect(err).To(BeNil(), "Should be able to create a Request")
			rj.SetRequest(request)
			//			runtime.Breakpoint()
			err = rj.Execute()
			Expect(err).To(BeNil(), " make a get request without problems")
			response, err := rj.GetResult()
			Expect(err).To(BeNil(), "Should be able to create a Request")
			defer response.Body.Close()
			cookies := response.Cookies()
			Expect(cookies[0].Name).To(Equal("times"), "the name of the cookie should be times")
			Expect(cookies[0].Value).To(Equal("1"), "the name of the cookie should be times")

		})
		FIt("Will save the cookies passed by the server for all the request", func() {

			rj, err := UglySpider.NewJobRequest(fakeServer.URL)
			Expect(err).To(BeNil(), "Should be able to create a Request")
			rj.SetRequest(request)
			//			runtime.Breakpoint()
			err = rj.Execute()
			err = rj.Execute()
			err = rj.Execute()
			err = rj.Execute()
			Expect(err).To(BeNil(), " make a get request without problems")
			res, err := rj.GetResult()
			Expect(err).To(BeNil(), " make a get request without problems")
			defer res.Body.Close()
			cookies := res.Cookies()
			Expect(cookies[0].Name).To(Equal("times"), "the name of the cookie should be times")
			Expect(cookies[0].Value).To(Equal("111"), "the is executed 3 times should have the cookies increase by 3")

		})

	})
})

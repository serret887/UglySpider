package coach_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/serret887/UglySpider/coach"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/ogle"
)

var _ = Describe("Coach", func() {
	Context("given an spcific matcher and a website coach return the ocurrences of the matcher", func() {
		var fakeServer *httptest.Server
		BeforeEach(func() {
			fakeServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cookie, err := r.Cookie("times")
				if err != nil {
					cookie = &http.Cookie{Name: "times"}
				}

				cookie.Value += "1"
				http.SetCookie(w, cookie)
				fmt.Fprintln(w)
			}))
		})

		It("Given a matcher for all <a> coach makethe request and return the result of the matchers ", func() {
			c, err := coach.NewCoach(1, "")
			Expect(err).To(BeNil(), "the error creating a acoach should be nil")
			// call process
			var urls []string
			urls = append(urls, fakeServer.URL)
			ogleChan := make(chan []*ogle.Ogle)
			c.Process(ogleChan, urls...)

		})
	})
})

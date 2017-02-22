package carScrapper_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/serret887/carScrapper"

	"bytes"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Worker", func() {
	Context("Worker execute any http request", func() {
		var response = "hello from the server"

		It("make an http get request", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(r)
				fmt.Fprint(w, response)
			}))
			defer ts.Close()

			w := &Worker{}
			w.SetURL(ts.URL)
			fmt.Println(ts.URL)
			var b bytes.Buffer
			//			runtime.Breakpoint()
			n, err := w.Write(&b)
			Expect(err).To(BeNil(), "there should not be error")
			Expect(n).To(Equal(len(response)), "The amount of data of the read")
			//		runtime.Breakpoint()
			Expect(b.String()).To(BeEquivalentTo(response), "the response of the write should be the same cause is getting the request via http")

		})

		It("if the URL is not set it return an empty buffer", func() {
			w := Worker{}
			w.SetURL("")
			var b bytes.Buffer
			res, err := w.Write(&b)
			Expect(res).To(BeNil())
			Expect(err).To(Equal(ErrMissingURL), "should flag the error cause there is no URL setted")
		})

		It("proxy all the request over the specifyed address", func() {
			w := &Worker{}
			w.SetProxyFromURL("http//:localhost:9754")
			var b bytes.Buffer
			w.Write(&b)
		})

		Context("Benchmarking making the requests", func() {

		})
	})

})

func readResourceFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return ioutil.ReadAll(f)
}

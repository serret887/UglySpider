package carScrapper

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

// errors

// ErrMissingURL this error is flaged if we dont set first the
// URL that we want to scrappe
var ErrMissingURL = errors.New("Before call Write youshould set first the URL you want to scrappe")

// Worker make a request to the URL setted before call the
// writer method. The advantage is we can use this interface
// and linked to any other method that use an interface
type Worker interface{}

// NewProxySpider create a new worker that proxy all their
// request over a proxy specifyed in the pAddress parameter
func NewProxySpider(pAddress string) (Worker, error) {
	p, err := url.Parse(pAddress)
	if err != nil {
		return nil, err
	}
	dialer, err := proxy.FromURL(p, proxy.Direct)
	ptransport := http.Transport
}

type worker struct {
	proxy, url *url.URL
}

// SetURL set the current URL for the worker to download
func (w *worker) SetURL(address string) error {
	var err error
	w.url, err = url.Parse(address)
	return err

}

// SetProxyFromURL set the proxy for the request
func (w *Worker) SetProxyFromURL(pAddress string) error {
	var err error
	w.proxy, err = url.Parse("pAddress")
	return err
}

func (w *Worker) Write(b io.Writer) (int, error) {
	res, err := http.Get(w.url.String())
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading the stream", err)
	}
	b.Write(result)
	// n1, err := io.Copy(b, result)
	fmt.Println("without proxy")
	return len(result), err

	// dialer, err := proxy.FromURL(w.proxy, proxy.Direct)
	// if err != nil {
	// 	return 0, err
	// }
	// transport := &http.Transport{Dial: dialer.Dial}
	// client := &http.Client{Transport: transport}

	// res, err := client.Get("http://check.torproject.org")
	// if err != nil {
	// 	return 0, err
	// }

	// defer res.Body.Close()
	// result, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal("Error reading the stream", err)
	// }
	// b.Write(result)
	// // n1, err := io.Copy(b, result)

	// return len(result), err
}

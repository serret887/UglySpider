package UglySpider

import (
	"log"
	"net/http"
	"time"
	"net/url"
 "golang.org/x/net/publicsuffix"
 "net/http/cookiejar"

)


// Worker make a request to the URL setted before call the
// writer method. The advantage is we can use this interface
// and linked to any other method that use an interface
type Worker interface{
	 Do(req *http.Request)(*http.Response,error)
}

// NewProxySpider create a new worker that proxy all their
// request over a proxy specifyed in the pAddress parameter
func NewProxySpider(pAddress string) (Worker, error) {
	pURL, err := url.Parse(pAddress)
	if err != nil {
		return nil, err
	}

// setting up the transport
	ptransport := &http.Transport{
IdleConnTimeout: 30*time.Second,
Proxy:http.ProxyURL(pURL),
	 }
	 //setting up COOKIES
	 jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	 if err != nil {
					 log.Fatal(err)
	 }
	 client := &http.Client{
		 Jar:jar,
		 Transport: ptransport,
	 }

	 // Creating the Worker
	 w := &worker{
		 Client: client,
	 }
return w , nil
}

type worker struct {
	*http.Client
}
//
// // SetURL set the current URL for the worker to download
// func (w *worker) SetURL(address string) error {
// 	var err error
// 	w.url, err = url.Parse(address)
// 	return err
//
// }
//
// // SetProxyFromURL set the proxy for the request
// func (w *Worker) SetProxyFromURL(pAddress string) error {
// 	var err error
// 	w.proxy, err = url.Parse("pAddress")
// 	return err
// }
//
// func (w *Worker) Write(b io.Writer) (int, error) {
// 	res, err := http.Get(w.url.String())
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	defer res.Body.Close()
// 	result, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal("Error reading the stream", err)
// 	}
// 	b.Write(result)
// 	// n1, err := io.Copy(b, result)
// 	fmt.Println("without proxy")
// 	return len(result), err
//
// 	// dialer, err := proxy.FromURL(w.proxy, proxy.Direct)
// 	// if err != nil {
// 	// 	return 0, err
// 	// }
// 	// transport := &http.Transport{Dial: dialer.Dial}
// 	// client := &http.Client{Transport: transport}
//
// 	// res, err := client.Get("http://check.torproject.org")
// 	// if err != nil {
// 	// 	return 0, err
// 	// }
//
// 	// defer res.Body.Close()
// 	// result, err := ioutil.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	log.Fatal("Error reading the stream", err)
// 	// }
// 	// b.Write(result)
// 	// // n1, err := io.Copy(b, result)
//
// 	// return len(result), err
// }

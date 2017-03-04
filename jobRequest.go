package UglySpider

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"runtime"

	"github.com/serret887/UglySpider/workerPool"

	"golang.org/x/net/publicsuffix"

	"time"

	"fmt"
)

type jobRequest struct {
	Transport *http.Transport
	*http.Client
	proxyURL url.URL
	Response *http.Response
	req      *http.Request
	err      error
}

type JobRequest interface {
	SetRequest(*http.Request)
	workerPool.Job
	GetResult() (*http.Response, error)
}

func NewSimpleJobRequest() (JobRequest, error) {
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}
	// Cookies
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	// Creating the Client
	client := &http.Client{
		Jar:       jar,
		Transport: tr,
	}

	// cerating the JOB REQ
	jobReq := &jobRequest{
		Transport: tr,

		Client: client,
	}

	// returning the JOB
	return jobReq, nil

}

func NewJobRequest(proxyAddr string) (JobRequest, error) {
	runtime.Breakpoint()
	pURL, err := url.Parse(proxyAddr)
	if err != nil {
		return nil, err
	}

	// creating transport
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
		Proxy:           http.ProxyURL(pURL),
	}

	// Cookies
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	// Creating the Client
	client := &http.Client{
		Jar:       jar,
		Transport: tr,
	}

	// cerating the JOB REQ
	jobReq := &jobRequest{
		Transport: tr,
		proxyURL:  *pURL,
		Client:    client,
	}

	// returning the JOB
	return jobReq, nil
}

func (jb *jobRequest) SetRequest(request *http.Request) {
	jb.req = request
}

func (jb *jobRequest) Execute() error {
	jb.Response, jb.err = jb.Do(jb.req)
	return jb.err
}

func (jb *jobRequest) Close() error {
	jb.Transport.CancelRequest(jb.req)
	jb.Transport.CloseIdleConnections()
	return nil
}

func (jb *jobRequest) GetResult() (*http.Response, error) {
	return jb.Response, jb.err
}

func (jb *jobRequest) String() string {
	return fmt.Sprintf("Client with proxy %v in the request %v", jb.proxyURL, jb.req)
}

// SetRequestMozillaHeaders create a request with the headers of the mozilla browser
func SetRequestMozillaHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch, br")
	req.Header.Set("Accept-Language", "en,es;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	//	req.Header.Set("Cookie", "csrftoken=m3sqBiLwPiEQhHRNYaiY4JbKWmEQXTwr; dwf_section_edit=False; dwf_sg_task_completion=False; _ga=GA1.2.154775868.1488591736; _gat=1")
	//	req.Header.Set("Host", "developer.mozilla.org")
	req.Header.Set("Referer", "//www.google.com/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	return req
}

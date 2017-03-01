package UglySpider

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

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

func NewJobRequest(proxyAddr string) (JobRequest, error) {

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

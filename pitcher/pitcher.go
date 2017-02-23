package pitcher

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Pitcher struct {
	pURL *url.URL
	*http.Client
}

// NewPitcher is the constructor of a new Pictcher the parameter
// passed to the constructor is the
func NewPitcher(pAddress string) (*Pitcher, error) {
	// if pAddress == "" {
	// 	client, err := newSimpleSpider()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &Pitcher{Client: client}, nil
	// }
	pURL, err := url.Parse(pAddress)
	if err != nil {
		return nil, err
	}
	client, err := NewProxySpider(pURL)
	if err != nil {
		return nil, err
	}
	return &Pitcher{pURL: pURL, Client: client}, nil
}

// NewProxySpider create a new worker that proxy all their
// request over a proxy specifyed in the pAddress parameter
func NewProxySpider(pURL *url.URL) (*http.Client, error) {

	// setting up the transport
	ptransport := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
		Proxy:           http.ProxyURL(pURL),
	}
	//setting up COOKIES
	client, err := newSimpleSpider()
	if err != nil {
		return nil, err
	}
	client.Transport = ptransport
	// Creating the Client
	return client, nil
}
func newSimpleSpider() (*http.Client, error) {
	//setting up COOKIES
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	return &http.Client{
		Jar: jar,
	}, nil

}

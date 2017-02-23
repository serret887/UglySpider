package pitcher

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

// NewProxySpider create a new worker that proxy all their
// request over a proxy specifyed in the pAddress parameter
func NewProxySpider(pAddress string) (*http.Client, error) {
	pURL, err := url.Parse(pAddress)
	if err != nil {
		return nil, err
	}

	// setting up the transport
	ptransport := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
		Proxy:           http.ProxyURL(pURL),
	}
	//setting up COOKIES
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Jar:       jar,
		Transport: ptransport,
	}

	// Creating the Worker
	return client, nil
}

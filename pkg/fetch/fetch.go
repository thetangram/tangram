package fetch

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// Request to fetch a view or component
type Request struct {
	*http.Request
	timeout time.Duration
}

// New request
func New(source *http.Request, timeout time.Duration, url string) (r *Request, err error) {
	temp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	r = &Request{
		Request: temp,
		timeout: timeout,
	}
	// TODO here add/rewrite the headers
	//      for example, X-Forwarded-Host, etc
	//newReq.Header.Add("If-None-Match", `W/"wyzzy"`)
	return
}

func NewSimple(timeout time.Duration, url string) (r *Request, err error) {
	temp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	r = &Request{
		Request: temp,
		timeout: timeout,
	}
	return
}

// Fetch a request and returns a DOM tree
func (r *Request) Fetch() (root *html.Node, err error) {
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
		Timeout: r.timeout,
	}
	response, err := client.Do(r.Request)
	if err != nil {
		// Error performing request
		log.Printf("Error fetching %v: %v\n", r.Request.URL, err)
		return
	}
	defer response.Body.Close()
	root, err = html.Parse(response.Body)
	return
}

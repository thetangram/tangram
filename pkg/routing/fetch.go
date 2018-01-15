package routing

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type Request struct {
	*http.Request
	timeout time.Duration
}

func new(source *http.Request, timeout time.Duration, url string) (r *Request, err error) {
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

func (r *Request) fetch() (root *html.Node, err error) {
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
		Timeout: r.timeout,
	}
	response, err := client.Do(r.Request)
	if err != nil {
		// Error performing request
		log.Printf("Error fetching %v: %v\n", r.Request.URL, err)
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	root, err = html.Parse(response.Body)
	return
}

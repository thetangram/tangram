package routing

import (
	"log"
	"net/http"

	"github.com/thetangram/tangram/pkg/conf"
	"golang.org/x/net/html"
)

type Router struct {
	conf.Route
}

func (r *Router) Register() {
	http.Handle(r.Path(), r)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	request, err := new(req, r.Timeout(), r.URL())
	if err != nil {
		// We cannot create the request. Log and return
		log.Printf("Error creating target request. target URL: %v. Error: %v\n", r.URL(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	root, err := request.fetch()
	if err != nil {
		// Error performing request
		log.Printf("Error fetching url %v. Error: %v\n", r.URL(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	html.Render(w, root)
}

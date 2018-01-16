package routing

import (
	"log"
	"net/http"

	"github.com/thetangram/tangram/pkg/composer"
	"github.com/thetangram/tangram/pkg/conf"
	"github.com/thetangram/tangram/pkg/fetch"
	"golang.org/x/net/html"
)

// Router is a routing configuration
type Router struct {
	conf.Route
}

// Register a route in the default HTTP server
func (r *Router) Register() {
	http.Handle(r.Path(), r)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	request, err := fetch.New(req, r.Timeout(), r.URL())
	if err != nil {
		// We cannot create the request. Log and return
		log.Printf("Error creating target request. target URL: %v. Error: %v\n", r.URL(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	root, err := request.Fetch()
	if err != nil {
		// Error performing request
		log.Printf("Error fetching url %v. Error: %v\n", r.URL(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if composedNode, err := composer.Compose(root); err == nil {
		html.Render(w, &composedNode)
	} else {
		// TODO see how to deal with composition error
		log.Printf("Error composing url %v. Error: %v\n", r.URL(), err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

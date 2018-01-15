package routing

import (
	"log"
	"net/http"

	"github.com/thetangram/tangram/pkg/conf"
    "github.com/thetangram/tangram/pkg/fetch"
    "github.com/thetangram/tangram/pkg/composer"
	"golang.org/x/net/html"
)

type Router struct {
	conf.Route
}

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

    if composedNode, err:=composer.Compose(root); err == nil {
        html.Render(w, composedNode)
    } else {
        // TODO ver como tratar el error de composicion!!!
        log.Printf("Error composing url %v. Error: %v\n", r.URL(), err)
        w.WriteHeader(http.StatusInternalServerError)
    }
}

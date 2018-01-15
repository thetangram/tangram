package routing

import (
	"fmt"
	"net/http"

	"github.com/thetangram/tangram/pkg/conf"
)

type Router struct {
	conf.Route
}

func (r *Router) Register() {
	http.Handle(r.Path(), r)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, r.URL())
}

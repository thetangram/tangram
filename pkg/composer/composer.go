package composer

import (
	"log"
	"time"

	"github.com/thetangram/tangram/pkg/fetch"
	"golang.org/x/net/html"
)

const (
	dataSrc           = "data-src"
	dataName          = "data-name"
	dataTimeout       = "data-timeout"
	dataHeaderFilter  = "data-headers-filter"
	dataCookiesFilter = "data-cookies-filter"
	dataTTL           = "data-ttl"
)

type holderAttributes struct {
	src           string
	name          string
	timeout       time.Duration
	headersFilter []string
	cookiesFilter []string
	ttl           time.Duration
}

// Compose a node
func Compose(root *html.Node) (node html.Node, err error) {
	node = processNode(root)
	return
}

func processNode(node *html.Node) html.Node {
	if node.Type == html.ElementNode {
		if isHolder, target := holder(node); isHolder == true {
			request, err := fetch.NewSimple(target.timeout, target.src) // TODO ver de donde sacar los datos
			if err != nil {
				// We cannot create the request. Log and return
				log.Printf("Error creating target request. target URL: %v. Error: %v\n", target.src, err)
				return *node
			}
			component, err := request.Fetch()
			if err == nil {
				clean(node)
				processed := processNode(component)
				node.AppendChild(&processed)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		processNode(c)
	}
	return *node
}

func holder(node *html.Node) (found bool, holderAttrs holderAttributes) {
	for _, a := range node.Attr {
		switch key := a.Key; key {
		case dataSrc:
			found = true
			holderAttrs.src = a.Val
		case dataName:
			holderAttrs.name = a.Val
			/*
			        // Here we should parse the tag value to concrete type (duration, array,...)
			        case dataTimeout:
						holderAttrs.timeout = a.Val
					case dataHeadersFilter:
						holderAttrs.headersFilter = a.Val
					case dataCookiesFilter:
						holderAttrs.cookiesFilter = a.Val
					case dataTTL:
						holderAttrs.ttl = a.Val
			*/
		}
	}
	return
}

func clean(node *html.Node) {
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode {
			node.RemoveChild(n)
		}
	}
	for _, a := range node.Attr {
		if a.Key == dataSrc {
			// TODO here remove this attribute from node.Attr array
			//fmt.Printf("    removing attribute %v\n", a.Key)
		}
	}
}

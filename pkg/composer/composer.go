package composer

import (
	"fmt"
	"log"
	"time"

	"github.com/thetangram/tangram/pkg/fetch"
	"golang.org/x/net/html"
)

const (
	dataLocationAttr = "data-src"
)

type holderAttributes struct {
	src          string
	name         string
	timeout      time.Duration
	headerFilter []string
	cookieFilter []string
	ttl          time.Duration
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

func holder(node *html.Node) (bool, holderAttributes) {
	for _, a := range node.Attr {
		if a.Key == dataLocationAttr {
			return true, holderAttributes{
				src:     a.Val,
				timeout: 5 * time.Second, // TODO Hardcoded!! Should come from attributes
				// TODO process the rest of attributes
			}
		}
	}
	return false, holderAttributes{}
}

func clean(node *html.Node) {
	fmt.Println("entering to clean node")
	fmt.Println("  removing child nodes...")
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode {
			node.RemoveChild(n)
		}
	}
	fmt.Println("  removing tangram attributes...")
	for _, a := range node.Attr {
		if a.Key == dataLocationAttr {
			// TODO here remove this attribute from node.Attr array
			fmt.Printf("    removing attribute %v\n", a.Key)
		}
	}
}

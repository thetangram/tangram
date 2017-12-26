package main

import (
	"fmt"
)

var (
	version   = "development"
	build     = "undefined"
	buildDate = "unknown"
)

func main() {
	fmt.Printf("Tangram test\n")
	fmt.Printf("version:   %s\n", version)
	fmt.Printf("build:     %s\n", build)
	fmt.Printf("buildDate: %s\n", buildDate)
}

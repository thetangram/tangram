package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	defaultAddress = ":8080"
)

var (
	version   = "development"
	build     = "undefined"
	buildDate = "unknown"
)

func healthy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}

func ready() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}

func alive() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}

func address() string {
	return defaultAddress
}

func main() {
	fmt.Printf("Tangram\n")
	fmt.Printf("  version:    %s\n", version)
	fmt.Printf("  build:      %s\n", build)
	fmt.Printf("  build date: %s\n", buildDate)

	http.HandleFunc("/healthy", healthy())
	http.HandleFunc("/ready", ready())
	http.HandleFunc("/alive", alive())
	log.Fatal(http.ListenAndServe(address(), nil))
}

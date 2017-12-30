package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultAddress = ":8080"
	// application exit status
	successExitStatus = 0
	sigTermExitStatus = 1
)

var (
	version   = "development"
	build     = "undefined"
	buildDate = "unknown"
	server    http.Server
)

// Liveness check.
// Used to verify if the application is running.
// An application is healthy if status code is >= 200 && <400
func healthy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}

// Readiness check.
// Used to verify if application is ready to serve client request.
// An application is ready if status code is >= 200 && <400
func ready() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}

func address() string {
	return defaultAddress
}

func shutdown() {
	fmt.Printf("Shutting down server...\n")
	if err := server.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	fmt.Printf("Tangram server stopped.\n")
	os.Exit(successExitStatus)
}

func main() {
	// print application info
	fmt.Printf("Tangram\n")
	fmt.Printf("  version:    %s\n", version)
	fmt.Printf("  build:      %s\n", build)
	fmt.Printf("  build date: %s\n", buildDate)

	// configure HTTP server and register application status entrypoints
	server := &http.Server{Addr: address()}
	http.HandleFunc("/healthy", healthy())
	http.HandleFunc("/ready", ready())
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	// deal with Ctrl+C (SIGTERM) and graceful shutdown
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	shutdown()
	os.Exit(sigTermExitStatus)
}

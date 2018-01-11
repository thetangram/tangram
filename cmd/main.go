// Package main provides de application entry point (main) and the
// procedures related to application life cycle.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cemugo/tangram/pkg/conf"
)

const (
	// application exit statuses
	successExitStatus            = 0
	errorLoadingConfig           = 1
	errorStopintServerStatusCode = 2
)

var (
	version   = "development"
	build     = "undefined"
	buildDate = "unknown"
	isReady   = false
)

func main() {
	printBanner()
	conf, err := conf.Load()
	if err != nil {
		log.Printf("Error loading configuration. Error: %v\n", err)
		os.Exit(errorLoadingConfig)
	}
	server := startHTTPServer(conf)
	waitAndShutdown(server, conf.ShutdownTimeout())
}

// print application info
func printBanner() {
	log.Println("Tangram")
	log.Printf("  version:      %s\n", version)
	log.Printf("  build:        %s\n", build)
	log.Printf("  build date:   %s\n", buildDate)
	log.Printf("  startup date: %s\n", time.Now().Format(time.RFC3339))
}

func startHTTPServer(c conf.Config) *http.Server {
	// configure HTTP server and register application status entrypoints
	server := &http.Server{
		Addr:         c.Address(),
		ReadTimeout:  c.ReadTimeout(),
		WriteTimeout: c.WriteTimeout(),
	}
	http.HandleFunc("/healthy", healthyHandler)
	http.HandleFunc("/ready", readyHandler)
	go func() {
		log.Printf("Listening on %s\n", c.Address())
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Cannot start HTTP server. Error: %s", err)
		}
	}()
	isReady = true
	return server

	//Test forking
}

func waitAndShutdown(server *http.Server, timeout time.Duration) {
	// deal with Ctrl+C (SIGTERM) and graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	isReady = false
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Printf("Shutting down, with a timeout of %s.\n", timeout)
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error stopping http server. Error: %v\n", err)
		os.Exit(errorStopintServerStatusCode)
	} else {
		log.Println("Tangram server stoped")
		os.Exit(successExitStatus)
	}
}

// Healthy check handler.
// Used to verify if the application is running.
// An application is healthy if its healthy status code is >= 200 && <400
func healthyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

// Readiness check.
// Used to verify if application is ready to serve client request.
// An application is ready if status code is >= 200 && <400
func readyHandler(w http.ResponseWriter, r *http.Request) {
	if isReady {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "NO")
	}
}

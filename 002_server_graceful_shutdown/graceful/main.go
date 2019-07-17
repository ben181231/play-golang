package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getHelloWorldHandlerFunc() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)

		time.Sleep(5 * time.Second)

		_, _ = rw.Write([]byte("Hello World"))
	}
}

func getShutdownHandlerFunc() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/shutdown", getShutdownHandlerFunc())
	mux.HandleFunc("/", getHelloWorldHandlerFunc())

	httpServer := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	doneChan := make(chan bool)
	go func(server *http.Server, done chan<- bool) {
		log.Printf("the server is running on PID %d.", os.Getpid())
		log.Print("the server is listening on port 8080.")

		if err := server.ListenAndServe(); err != nil {
			log.Printf("Error: %s", err)
		}

		done <- true
	}(httpServer, doneChan)

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signalChan:
	case <-doneChan:
	}

	log.Print("the server is shutting down gracefully")
	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown http server: %s", err.Error())
	}
}

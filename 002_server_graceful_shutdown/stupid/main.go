package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func getHelloWorldHandlerFunc() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)

		time.Sleep(10 * time.Second)

		_, _ = rw.Write([]byte("Hello World"))
	}
}

func getShutdownHandlerFunc() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	}
}

func doSomethingElse() {
	log.Print("the server is shutdown gracefully")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/shutdown", getShutdownHandlerFunc())
	mux.HandleFunc("/", getHelloWorldHandlerFunc())

	log.Printf("the server is running on PID %d.", os.Getpid())
	log.Print("the server is listening on port 8080.")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("Error: %s", err)
	}

	doSomethingElse()
}

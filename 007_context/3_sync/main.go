package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ben181231/play-golang/007_context/3_sync/handler"
)

const (
	listenPort = 8080
)

func main() {
	addr := fmt.Sprintf("0.0.0.0:%d", listenPort)
	log.Printf("Server is listening on %s", addr)

	mux := http.NewServeMux()
	mux.Handle(
		"/timeout",
		handler.NewTimeoutHandler(time.Second, 2*time.Second),
	)
	mux.Handle("/", handler.NewLongRunningHandler(2*time.Second))

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Printf("Error: %s", err.Error())
	}
}

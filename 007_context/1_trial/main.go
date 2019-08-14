package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	listenPort = 8080
)

func getHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			rw.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintln(rw, "Hello World")
		}()

		ctx := r.Context()
		log.Printf("Get request with ctx: %v", ctx)

		if svc, ok := ctx.Value(http.ServerContextKey).(*http.Server); ok {
			log.Printf("Server on %s is handling the request", svc.Addr)
		}
	})
}

func main() {
	addr := fmt.Sprintf("0.0.0.0:%d", listenPort)
	log.Printf("Server is listening on %s", addr)
	if err := http.ListenAndServe(addr, getHandler()); err != nil {
		log.Printf("Error: %s", err.Error())
	}
}

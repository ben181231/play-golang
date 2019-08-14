package main

import (
	"fmt"
	"log"
	"net/http"

	"localhost/ben181231/007_context/2_values/requestid"
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
		rid := requestid.GetRequestID(ctx)
		log.Printf("Got request with ID: %s", rid)

	})
}

func main() {
	addr := fmt.Sprintf("0.0.0.0:%d", listenPort)
	log.Printf("Server is listening on %s", addr)

	handler := getHandler()
	handler = requestid.RequestIDAdapter{}.Adapt(handler)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Printf("Error: %s", err.Error())
	}
}

package main

import (
	"log"
	"net/http"

	"localhost/ben181231/006_goroutine_leak/handler/check"
	"localhost/ben181231/006_goroutine_leak/handler/leak"
	"localhost/ben181231/006_goroutine_leak/handler/no_leak"
)

func main() {
	/**
	 *
	 * This implementation is inspired by the following blog post:
	 * https://blog.minio.io/debugging-go-routine-leaks-a1220142d32c
	 *
	 */
	http.Handle("/leak/sum", leak.GetSumHandler())
	http.Handle("/no-leak/sum", noleak.GetSumHandler())
	http.Handle("/_count", check.GetGoroutinesCountHandler())
	http.Handle("/_stack", check.GetStackTraceHandler())

	log.Printf("Server is listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package check

import (
	"net/http"
	"runtime"
	"runtime/pprof"
	"strconv"
)

// GetGoroutinesCountHandler returns a handler telling number of goroutine running
func GetGoroutinesCountHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := runtime.NumGoroutine()
		w.Write([]byte(strconv.Itoa(count)))
	})
}

// GetStackTraceHandler returns a handler printing the stack trace of running goroutines
func GetStackTraceHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pprof.Lookup("goroutine").WriteTo(w, 2)
	})
}

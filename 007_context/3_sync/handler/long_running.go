package handler

import (
	"fmt"
	"net/http"
	"time"

	longrunning "github.com/ben181231/play-golang/007_context/3_sync/long_running"
)

func NewLongRunningHandler(d time.Duration) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			rw.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintln(rw, "Hello World")
		}()

		longrunning.Run(r.Context(), d)
	})
}

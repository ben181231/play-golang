package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	longrunning "github.com/ben181231/play-golang/007_context/3_sync/long_running"
)

func NewTimeoutHandler(timeout, duration time.Duration) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			rw.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintln(rw, "Hello World")
		}()

		timeoutCtx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		longrunning.Run(timeoutCtx, duration)
	})
}

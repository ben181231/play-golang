package requestid

import (
	"net/http"

	"github.com/google/uuid"
)

type RequestIDAdapter struct{}

func (RequestIDAdapter) Adapt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if rid := GetRequestID(ctx); rid == "" {
			rid = uuid.New().String()
			r = r.WithContext(WithRequestID(ctx, rid))
		}

		next.ServeHTTP(rw, r)
	})
}

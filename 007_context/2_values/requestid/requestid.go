package requestid

import "context"

var requestIDContextKey = struct{}{}

type requestID struct {
	value string
}

func GetRequestID(ctx context.Context) string {
	if rid, got := ctx.Value(requestIDContextKey).(*requestID); got {
		return rid.value
	}

	return ""
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(
		ctx,
		requestIDContextKey,
		&requestID{value: id},
	)
}

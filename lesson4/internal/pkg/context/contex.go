package context

import "context"

type contextRequestIDKey struct{}

func SetRequestID(parent context.Context, requestID string) context.Context {
	return context.WithValue(parent, contextRequestIDKey{}, requestID)
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(contextRequestIDKey{}).(string)
	if !ok {
		panic("can't get request id from context")
	}

	return requestID
}

package helper

import (
	"context"

	"github.com/go-chi/chi/middleware"
)

type ctxKeyRequestID int

// RequestIDKey is the key that holds the unique request ID in a request context.
const RequestIDKey ctxKeyRequestID = 0

func GetReqID(ctx context.Context) string {
	return middleware.GetReqID(ctx)
}

package middleware

type (
	ContextKey string
)

var (
	CONTEXT_CLAIM_KEY ContextKey = "claims"
	CONTEXT_JWT_KEY   ContextKey = "jwt"
	CONTEXT_CACHE_KEY ContextKey = "cache"
)

const (
	HEADER_CLIENT_ID        = "x-client-id"
	HEADER_CLIENT_SIGNATURE = "x-signature"
	HEADER_TIMESTAMP        = "x-timestamp"
)

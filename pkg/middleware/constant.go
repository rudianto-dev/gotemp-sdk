package middleware

type (
	ContextKey string
)

var (
	CONTEXT_CLAIM_KEY ContextKey = "claims"
	CONTEXT_JWT_KEY   ContextKey = "jwt"
	CONTEXT_CACHE_KEY ContextKey = "cache"
)

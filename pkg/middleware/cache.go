package middleware

import (
	"context"
	"net/http"

	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
)

func SetCache(cache *cache.DataSource) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if cache != nil {
				ctx = context.WithValue(ctx, CONTEXT_CACHE_KEY, cache)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetCache(c context.Context) cache.Service {
	return c.Value(CONTEXT_CACHE_KEY).(*cache.DataSource)
}

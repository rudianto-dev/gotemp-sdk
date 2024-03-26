package middleware

import (
	"context"
	"net/http"

	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
)

func JWT(j *token.JWT) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if j != nil {
				ctx = context.WithValue(r.Context(), CONTEXT_JWT_KEY, j)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetJWT(c context.Context) token.IJWTToken {
	return c.Value(CONTEXT_JWT_KEY).(*token.JWT)
}

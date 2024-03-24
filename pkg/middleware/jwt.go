package middleware

import (
	"context"
	"net/http"

	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
)

const (
	JWTContextKey = "jwt"
)

func JWT(jw *token.JWT) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if jw != nil {
				ctx = context.WithValue(r.Context(), JWTContextKey, jw)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetJWT(c context.Context) token.IJWTToken {
	return c.Value(JWTContextKey).(*token.JWT)
}

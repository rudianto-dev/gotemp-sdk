package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
)

type (
	TokenRetrieval func(r *http.Request) string
	ContextKey     string
)

var (
	CONTEXT_CLAIM_KEY ContextKey = "claims"
)

func TokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

func introspect(r *http.Request, tr ...TokenRetrieval) (payload token.Payload, err error) {
	var tokenString string
	for _, fn := range tr {
		if tokenString = fn(r); tokenString != "" {
			break
		}
	}
	if tokenString == "" {
		err = token.ErrNoTokenFound
		return
	}
	return GetJWT(r.Context()).Validate(tokenString)
}

func GuardAuthenticated(tr ...TokenRetrieval) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			payload, err := introspect(r, tr...)
			if err != nil {
				MakeLogEntry("GuardAuthentication - Introspect").Error(err)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// redis := jwe.NewDataRedis(GetRedis(r.Context()), GetLogEntry(r))
			// if token, err := redis.Find(r.Context(), pub.Subject, merchant, pri.DeviceId); err != nil || token != TokenFromHeader(r) {
			// 	MakeLogEntry("GuardAuthentication - Redis").Error(err)
			// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			// 	return
			// }
			bundle := &token.Payload{ID: payload.ID, UserID: payload.UserID, RoleType: payload.RoleType}
			ctx := context.WithValue(r.Context(), CONTEXT_CLAIM_KEY, bundle)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetClaims(r *http.Request) (c *token.Payload) {
	if value := r.Context().Value(CONTEXT_CLAIM_KEY); value != nil {
		c = value.(*token.Payload)
	}
	return
}

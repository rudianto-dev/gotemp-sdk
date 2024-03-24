package JWTToken

import "time"

type IJWTToken interface {
	Create(ttl time.Duration, content interface{}) (string, error)
	Validate(token string) (interface{}, error)
}

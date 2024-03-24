package token

type IJWTToken interface {
	Create(content Payload) (token string, expiredAt int64, err error)
	Validate(token string) (payload Payload, err error)
}

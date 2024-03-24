package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JWT) Create(content Payload) (token string, expiredAt int64, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.PrivateKey)
	if err != nil {
		j.Logger.Errorf("create: parse key: %w", err)
		return
	}

	now := time.Now().UTC()
	expiredAt = now.Add(j.TTL).Unix()

	claims := make(jwt.MapClaims)
	claims["dat"] = content    // Our custom data.
	claims["exp"] = expiredAt  // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix() // The time at which the token was issued.
	claims["nbf"] = now.Unix() // The time before which the token must be disregarded.

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		j.Logger.Errorf("create: parse key: %w", err)
		return
	}
	return
}

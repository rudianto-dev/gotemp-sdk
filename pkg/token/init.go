package token

import (
	"time"

	"errors"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type JWT struct {
	PrivateKey []byte
	PublicKey  []byte
	TTL        time.Duration
	Logger     *logger.Logger
}

var (
	ErrInvalidScope = errors.New("invalid scope in claims")
	ErrNoTokenFound = errors.New("auth: no credentials attached in request")
)

func New(privateKey []byte, publicKey []byte, ttl time.Duration, logger *logger.Logger) *JWT {
	return &JWT{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		TTL:        ttl,
		Logger:     logger,
	}
}

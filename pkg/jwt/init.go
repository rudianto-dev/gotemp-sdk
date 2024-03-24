package JWTToken

type JWT struct {
	privateKey []byte
	publicKey  []byte
}

func New(privateKey []byte, publicKey []byte) IJWTToken {
	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

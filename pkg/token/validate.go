package token

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JWT) Validate(token string) (payload Payload, err error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.PublicKey)
	if err != nil {
		j.Logger.Errorf("validate: parse key: %w", err)
		return
	}

	validate, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return
	}

	claims, ok := validate.Claims.(jwt.MapClaims)
	if !ok || !validate.Valid {
		err = fmt.Errorf("validate: invalid")
		return
	}
	payload = j.CollectPayload(claims["dat"])
	return
}

func (j *JWT) CollectPayload(data interface{}) Payload {
	m := data.(map[string]interface{})
	customer := Payload{}
	if ID, ok := m["id"].(string); ok {
		customer.ID = ID
	}
	if userID, ok := m["user_id"].(string); ok {
		customer.UserID = userID
	}
	if roleType, ok := m["role_type"].(string); ok {
		i, err := strconv.Atoi(roleType)
		if err == nil {
			customer.RoleType = int8(i)
		} else {
			customer.RoleType = 0
		}
	}
	return customer
}

package tool

import (
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func ParseToken(token, secretKey string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

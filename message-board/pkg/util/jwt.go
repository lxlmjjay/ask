package util

import (
	"message-board/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

)

var jwtSecret = []byte(setting.JwtSecret)
type Claims struct {
	UserId int `json:"user_id"`
	UserPass int `json:"user_pass"`
	jwt.StandardClaims
}

func GenerateToken(userId, userPass int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		userId,
		userPass,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "message-board",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
	func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

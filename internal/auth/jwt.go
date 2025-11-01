package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJwt(secret []byte, userId int, expirationSeconds int) (string, error) {
	expiration := time.Second * time.Duration(expirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   strconv.Itoa(userId),
		"expireAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

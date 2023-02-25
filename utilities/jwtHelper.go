package utilities

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("secret")

type MyCustomClaims struct {
	DataToSign string `json:"dataToSign"`
	jwt.StandardClaims
}

func GenerateAccessToken(dataToSign string) (string, error) {
	accessTokenClaims := MyCustomClaims{
		DataToSign: dataToSign,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = accessTokenClaims
	accessTokenString, err := token.SignedString(SecretKey)

	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func GenerateRefreshToken(dataToSign string) (string, error) {
	refreshTokenClaims := MyCustomClaims{
		DataToSign: dataToSign,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = refreshTokenClaims
	refreshTokenString, err := token.SignedString(SecretKey)

	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}

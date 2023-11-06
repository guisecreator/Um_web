package token

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var SecretKey = []byte("secret")

func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	expireUnix := time.Now().AddDate(1, 0, 0).Unix()
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["expired"] = expireUnix

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}

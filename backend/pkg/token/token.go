package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Token struct {
	value       string
	claims      Claims
	expiredTime time.Time
	RequestInfo RequestInfo
	rsa.PrivateKey
}

type RequestInfo struct {
	Method  string
	Host    string
	URL     http.Request
	Headers http.Header
}

type Claims struct {
	Id      int
	Expired int64
}

func (t *Token) GenerateToken(id int) (Token, error) {
	token := Token{
		RequestInfo: RequestInfo{
			Method:  "GET",
			Host:    "localhost",
			URL:     http.Request{},
			Headers: http.Header{},
		},
	}

	token.expiredTime = time.Now().AddDate(1, 0, 0)

	claims := Claims{
		Id: id,
		Expired: token.
			expiredTime.
			Unix(),
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	claimsBytes, errClaims := json.Marshal(claims)
	if errClaims != nil {
		return token, fmt.Errorf(
			"error encoding claims: %w", err)
	}

	encryptedBytes, err := Encrypt(
		sha256.New(),
		rand.Reader,
		&privateKey.PublicKey,
		claimsBytes,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	tokenStr := base64.StdEncoding.EncodeToString(encryptedBytes)

	token = Token{
		value:       tokenStr,
		expiredTime: token.expiredTime,
		PrivateKey:  *privateKey,
	}

	return token, nil
}

package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	Id      string
	Expired int64
}

func GenerateToken(id string) (Token, error) {
	t := Token{}
	expireUnix := time.Now().AddDate(1, 0, 0)
	claims := Claims{
		Id:      id,
		Expired: expireUnix.Unix(),
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	claimsBytes, errClaims := json.Marshal(claims)
	if errClaims != nil {
		return t, fmt.Errorf("error encoding claims: %w", err)
	}

	encryptedBytes, err := EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&privateKey.PublicKey,
		claimsBytes,
		nil,
	)
	if err != nil {
		panic(err)
	}

	tokenStr := base64.StdEncoding.EncodeToString(encryptedBytes)

	t = Token{
		value:      tokenStr,
		PrivateKey: *privateKey,
	}

	return t, nil
}

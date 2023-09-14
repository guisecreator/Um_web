package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
)

func Encrypt(hash hash.Hash, random io.Reader,
	public *rsa.PublicKey, msg []byte,
	label []byte) ([]byte, error) {

	msgLen := len(msg)
	step := public.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash,
			random, public,
			msg[start:finish],
			label)
		if err != nil {
			return nil, err
		}
		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}

	return encryptedBytes, nil
}

func Decrypt(hash hash.Hash, random io.Reader,
	private *rsa.PrivateKey, msg []byte,
	label []byte) ([]byte, error) {

	msgLen := len(msg)
	step := private.PublicKey.Size()
	var decryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := rsa.DecryptOAEP(
			hash, random,
			private, msg[start:finish],
			label)

		if err != nil {
			return nil, err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return decryptedBytes, nil
}

func DecryptBytes(code string,
	key rsa.PrivateKey) (Claims, error) {
	var claims Claims
	b, errDecode := base64.StdEncoding.DecodeString(code)
	if errDecode != nil {
		return Claims{}, errDecode
	}

	decrypted, errDecrypt := Decrypt(sha256.New(),
		rand.Reader, &key, b, nil)
	if errDecrypt != nil {
		return Claims{}, errDecrypt
	}

	errUnmarshal := json.Unmarshal(decrypted, &claims)
	if errUnmarshal != nil {
		return Claims{}, errUnmarshal
	}
	return claims, nil
}

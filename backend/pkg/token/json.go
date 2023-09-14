package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/guisecreator/um_backend/graphql/model"
)

func (t *Token) PutTokenInJSON(token string) (string, error) {
	tokenErr := errors.New("error marshalling token")

	token = t.value
	TokenINFO := model.AuthInfo{
		Token: token,
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		return tokenErr.Error(), err
	}

	token = base64.
		StdEncoding.
		EncodeToString(jsonToken)
	TokenINFO.Token = token

	return token, nil
}

func PullTokenFromJSON(token string) (string, error) {
	tokenErr := errors.New("invalid token")

	decodedToken, err := base64.
		StdEncoding.
		DecodeString(token)
	if err != nil {
		return "", tokenErr
	}

	decodedTokenString := decodedToken[:len(decodedToken)-1]
	tokenInfo := model.AuthInfo{
		Token: token,
	}

	unmarshalTokenErr := json.Unmarshal(
		decodedTokenString,
		&tokenInfo)
	if unmarshalTokenErr != nil {
		panic(unmarshalTokenErr)
		return "", tokenErr
	}

	return token, nil
}

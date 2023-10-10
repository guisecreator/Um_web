package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/guisecreator/um_web/graphql/model"
)

func (t *Token) PutTokenInJSON() (string, error) {
	tokenErr := errors.New("error marshalling token")

	token := Token{
		value: t.value,
	}

	//token = t.value
	TokenINFO := model.AuthInfo{
		Token: token.value,
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		return tokenErr.Error(), err
	}

	token.value = base64.
		StdEncoding.
		EncodeToString(jsonToken)
	TokenINFO.Token = token.value

	return token.value, nil
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

func (t *Token) ToJsonString() (string, error) {
	tokenStr := t.value

	authInfo := &model.AuthInfo{Token: tokenStr}
	b, errJson := json.Marshal(authInfo)
	if errJson != nil {
		return "", errJson
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

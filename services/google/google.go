package google

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type GoogleCredential struct {
	jwt.Claims
	Sub     string `json:"sub"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func ParseJWTCredential(token string) (*GoogleCredential, error) {

	// TODO: Parse token with Google's Public Certificate Key
	t, _, err := jwt.NewParser().ParseUnverified(token, &GoogleCredential{})
	if err != nil {
		return nil, err
	}

	if c, ok := t.Claims.(*GoogleCredential); ok {
		return c, nil
	}
	return nil, errors.New("could not cast into *GoogleCredential")
}

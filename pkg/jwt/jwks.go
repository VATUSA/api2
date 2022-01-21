package jwt

import (
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var KeySet jwk.Set

func ParseKeyset(keyset string) error {
	var err error
	KeySet, err = jwk.ParseString(keyset)
	if err != nil {
		return err
	}
	return nil
}

func ParseToken(token string) (jwt.Token, error) {
	t, err := jwt.Parse([]byte(token), jwt.WithKeySet(KeySet), jwt.WithValidate(true))

	return t, err
}

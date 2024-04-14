package validator

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrNotValid = errors.New("provided token is not valid")
	ErrExpired  = errors.New("provided token is not valid")
	ErrNotAdmin = errors.New("token does not belongs to admin")
)

func CheckAdmin(token string) error {
	key := fetchSecretKey()
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if ok && claims["admin"] == true {
		return nil
	} else {
		return ErrNotAdmin
	}
}

func Authorize(token string) error {
	key := fetchSecretKey()
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	var ve *jwt.ValidationError
	if errors.As(err, &ve) {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return ErrExpired
		}
	}

	if err != nil {
		return err
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if ok && t.Valid {
		return nil
	} else {
		return ErrNotValid
	}
}

func fetchSecretKey() string {
	const key = "SECRET_KEY"

	if v := os.Getenv(key); v != "" {
		return v
	}

	return "guulik"
}

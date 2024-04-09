package validator

import (
	"Avito_trainee_assignment/internal/config/constants"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	ErrNotValid = errors.New("provided token is not valid")
	ErrNotAdmin = errors.New("token does not belongs to admin")
	//ErrNotUser  = errors.New("token does not belongs to user")
)

func CheckAdmin(token string) error {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
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
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})
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

package validator

import "errors"

var (
	ErrNotValid = errors.New("provided token is not valid")
	ErrNotAdmin = errors.New("token does not belongs to admin")
	ErrNotUser  = errors.New("token does not belongs to user")
)

func CheckAdmin(token string) error {
	//TODO implement me
	if !valid(token) {
		return ErrNotValid
	}
	//if admin
	return nil
	//else
	return ErrNotAdmin
}

func CheckUser(token string) error {
	//TODO implement me
	if !valid(token) {
		return ErrNotValid
	}
	//if user
	return nil
	//else
	return ErrNotUser
}

func valid(token string) bool {
	//TODO implement me
	return true
	return false
}

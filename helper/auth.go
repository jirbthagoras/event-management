package helper

import (
	"errors"
)

func ComparePassword(password string, requestPassword string) error {
	if password != requestPassword {
		return errors.New("password not match")
	}

	return nil
}

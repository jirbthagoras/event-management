package helper

import (
	"errors"
)

func ComparePassword(password string, requestPassword string) error {
	if password != requestPassword {
		return errors.New("password did not not match")
	}
	return nil
}

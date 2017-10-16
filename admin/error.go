package admin

import (
	"errors"
)

const adminExist = "adminExist"
const adminNotExist = "adminNotExist"
const invalidPassword = "invalidPassword"

func NewAdminExistError() error {
	return errors.New(adminExist)
}

func NewAdminNotExistError() error {
	return errors.New(adminNotExist)
}

func NewInvalidPassword() error {
	return errors.New(invalidPassword)
}

func IsAdminExist(err error) bool {
	return err.Error() == adminExist
}

func IsAdminNotExist(err error) bool {
	return err.Error() == adminNotExist
}

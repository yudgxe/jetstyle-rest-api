package model

import (
	"errors"
)

var errMinPasswordLenght = errors.New("insufficient password length. Minimum of 8 characters")

const minPasswordLenght = 8

type User struct {
	ID             int32  `json:"id"`
	Login          string `json:"login"`
	Password       string `json:"password,omitempty"`
	HashedPassword string `json:"-"`
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) Validate() error {
	if len(u.Password) < minPasswordLenght {
		return errMinPasswordLenght
	}

	return nil
}

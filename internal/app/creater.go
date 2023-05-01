package app

import (
	"context"
	"log"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type Creater struct {
	store store.Store
}

func NewCreator(store store.Store) *Creater {
	return &Creater{
		store: store,
	}
}

func (cr *Creater) CreateUser(login, password string) error {
	user := &model.User{
		Login:    login,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.HashedPassword = string(passwordHash)

	if err := cr.store.User().Create(context.Background(), user); err != nil {
		return err
	}

	log.Printf("User created his id = %d", user.ID)

	return nil
}

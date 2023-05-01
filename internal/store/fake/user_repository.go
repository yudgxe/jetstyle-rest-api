package fake

import (
	"context"
	"errors"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
)

type UserRepository struct {
	users map[string]*model.User
}

var _ store.UserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*model.User),
	}
}

func (ur *UserRepository) Create(ctx context.Context, u *model.User) error {
	u.ID = int32((len(ur.users) + 1))
	ur.users[u.Login] = u

	return nil
}

func (ur *UserRepository) FindByLogin(ctx context.Context, login string) (*model.User, error) {
	user, ok := ur.users[login]
	if !ok {
		return nil, errors.New("no users with this login")
	}

	return user, nil
}

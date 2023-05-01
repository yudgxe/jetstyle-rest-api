package store

import (
	"context"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	FindByLogin(ctx context.Context, login string) (*model.User, error)
}

type TaskRepository interface {
	Create(ctx context.Context, t *model.Task) error
	FindById(ctx context.Context, id int32) (*model.Task, error)
	DeleteById(ctx context.Context, id int32) error
	UpdateById(ctx context.Context, id int32, t *model.Task) error
}

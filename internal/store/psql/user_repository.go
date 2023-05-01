package psql

import (
	"context"
	"database/sql"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
)

type UserRepository struct {
	db *sql.DB
}

var _ store.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func(ur * UserRepository) Create(ctx context.Context, u * model.User) error {
	if err := ur.db.QueryRowContext(ctx,
		`INSERT INTO users(
			login,
			password
		) VALUES ($1, $2) RETURNING id`,
		u.Login,
		u.HashedPassword,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func(ur *UserRepository) FindByLogin(ctx context.Context, login string) (*model.User,error) {
	user := &model.User{}

	if err := ur.db.QueryRowContext(ctx,
		`SELECT id, login, password FROM users WHERE login = $1`,
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.HashedPassword,
	); err != nil {
		return nil, err
	}

	return user, nil
}

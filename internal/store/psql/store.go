package psql

import (
	"database/sql"

	"github.com/yudgxe/jetstyle-rest-api/internal/store"


)

type Store struct {
	userRepository *UserRepository
	taskReposityry *TaskRepository
}

var _ store.Store = (*Store)(nil)

func New(db *sql.DB) *Store {
	return &Store{
		userRepository: NewUserRepository(db),
		taskReposityry: NewTaskRepository(db),
	}
}

func (s *Store) User() store.UserRepository {
	return s.userRepository
}

func (s *Store) Task() store.TaskRepository {
	return s.taskReposityry
}

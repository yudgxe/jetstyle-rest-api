package fake

import "github.com/yudgxe/jetstyle-rest-api/internal/store"

type Store struct {
	user *UserRepository
	task *TaskRepository
}

var _ store.Store = (*Store)(nil)

func NewStore() *Store {
	return &Store{
		user: NewUserRepository(),
		task: NewTaskRepository(),
	}
}

func (s *Store) User() store.UserRepository {
	return s.user
}

func (s *Store) Task() store.TaskRepository {
	return s.task
}

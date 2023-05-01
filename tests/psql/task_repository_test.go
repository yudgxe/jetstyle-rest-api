package psql_test

import (
	"context"
	"testing"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store/psql"
)

func TestTaskRepository_Create(t *testing.T) {
	s, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users", "tasks")

	user := model.GetTestUser(t)
	if err := s.User().Create(context.Background(), user); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	task := model.GetTestTask(t)
	task.Owner = user.ID

	if err := s.Task().Create(context.Background(), task); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestTaskRepository_GetById(t *testing.T) {
	s, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users", "tasks")

	user := model.GetTestUser(t)
	if err := s.User().Create(context.Background(), user); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	task := model.GetTestTask(t)
	task.Owner = user.ID

	if err := s.Task().Create(context.Background(), task); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	_, err := s.Task().FindById(context.Background(), task.ID)
	if err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestTaskRepository_DeleteById(t *testing.T) {
	s, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users", "tasks")

	user := model.GetTestUser(t)
	if err := s.User().Create(context.Background(), user); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	task := model.GetTestTask(t)
	task.Owner = user.ID

	if err := s.Task().Create(context.Background(), task); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	if err := s.Task().DeleteById(context.Background(), task.ID); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

package psql_test

import (
	"context"
	"testing"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store/psql"
)

func TestUserRepository_Create(t *testing.T) {
	s, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users")

	if err := s.User().Create(context.Background(), model.GetTestUser(t)); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestUserRepository_GetByLogin(t *testing.T) {
	s, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users")

	user := model.GetTestUser(t)
	if err := s.User().Create(context.Background(), user); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	_, err := s.User().FindByLogin(context.Background(), user.Login)
	if err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

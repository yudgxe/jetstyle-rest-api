package model

import "testing"

func GetTestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Login:    "login",
		Password: "password",
	}
}

func GetTestTask(t *testing.T) *Task {
	t.Helper()

	return &Task{
		Owner: 1,
		Name:  "name",
	}
}

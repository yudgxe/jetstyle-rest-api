package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/yudgxe/jetstyle-rest-api/internal/app"
	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store/fake"
)

func TestTask_HandleGet(t *testing.T) {
	user := model.GetTestUser(t)
	store := fake.NewStore()
	server := app.NewServer(store)
	app.NewCreator(store).CreateUser(user.Login, user.Password)

	task := model.GetTestTask(t)
	if err := store.Task().Create(context.Background(), task); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	tests := []struct {
		name         string
		id           interface{}
		expectedCode int
	}{
		{
			name:         "valid",
			id:           strconv.Itoa(int(task.ID)),
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid id",
			id:           strconv.Itoa(int(task.ID) + 1),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "invalid type",
			id:           "invalid",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "empty id",
			id:           "",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/tasks/"+test.id.(string), nil)

			req.SetBasicAuth(user.Login, user.Password)
			server.ServeHTTP(rec, req)
			if rec.Code != test.expectedCode {
				t.Errorf("\nGot error -> code was expected %d code was received %d", test.expectedCode, rec.Code)
			}
		})
	}
}
func TestTask_HandleCreate(t *testing.T) {
	user := model.GetTestUser(t)
	store := fake.NewStore()
	server := app.NewServer(store)
	app.NewCreator(store).CreateUser(user.Login, user.Password)

	tests := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"owner": 1,
				"name":  "name",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid params",
			payload: map[string]interface{}{
				"owner": 1,
				"name":  "nam",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "ivalid payload",
			payload:      "ivalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(test.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/tasks/", b)
			req.SetBasicAuth(user.Login, user.Password)
			server.ServeHTTP(rec, req)
			if rec.Code != test.expectedCode {
				t.Errorf("\nGot error in test: %s -> code was expected %d code was received %d", test.name, test.expectedCode, rec.Code)
			}
		})
	}
}

func TestTask_HandleDelete(t *testing.T) {
	user := model.GetTestUser(t)
	store := fake.NewStore()
	server := app.NewServer(store)
	app.NewCreator(store).CreateUser(user.Login, user.Password)

	task := model.GetTestTask(t)
	store.Task().Create(context.Background(), task)

	tests := []struct {
		name         string
		id           interface{}
		expectedCode int
	}{

		{
			name:         "valid",
			id:           strconv.Itoa(int(task.ID)),
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid id",
			id:           strconv.Itoa(int(task.ID) + 1),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "invalid type",
			id:           "invalid",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+test.id.(string), nil)

			req.SetBasicAuth(user.Login, user.Password)
			server.ServeHTTP(rec, req)
			if rec.Code != test.expectedCode {
				t.Errorf("\nGot error in test: %s -> code was expected %d code was received %d", test.name, test.expectedCode, rec.Code)
			}
		})
	}
}

func TestTask_HandlePut(t *testing.T) {
	user := model.GetTestUser(t)
	store := fake.NewStore()
	server := app.NewServer(store)
	app.NewCreator(store).CreateUser(user.Login, user.Password)

	task := model.GetTestTask(t)
	store.Task().Create(context.Background(), task)

	tests := []struct {
		name         string
		id           interface{}
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			id:   strconv.Itoa(int(task.ID)),
			payload: map[string]interface{}{
				"owner": 1,
				"name":  "name",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid params",
			id:   strconv.Itoa(int(task.ID)),
			payload: map[string]interface{}{
				"owner": 1,
				"name":  "nam",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid payload",
			id:           strconv.Itoa(int(task.ID)),
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid type id",
			id:   "invalid",
			payload: map[string]interface{}{
				"owner": 1,
				"name":  "name",
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(test.payload)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/tasks/"+test.id.(string), b)

			req.SetBasicAuth(user.Login, user.Password)
			server.ServeHTTP(rec, req)
			if rec.Code != test.expectedCode {
				t.Errorf("\nGot error in test: %s -> code was expected %d code was received %d", test.name, test.expectedCode, rec.Code)
			}
		})
	}
}

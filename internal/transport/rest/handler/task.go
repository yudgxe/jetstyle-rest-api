package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
)

type TaskHandler struct {
	store store.Store
}

func NewTaskHandler(s store.Store) *TaskHandler {
	return &TaskHandler{
		store: s,
	}
}

func (th *TaskHandler) Bind(mux *mux.Router) {
	mux.Handle("/", th.handleCreate()).Methods("POST")
	mux.Handle("/{id:[0-9]+}", th.handleGet()).Methods("GET")
	mux.Handle("/{id:[0-9]+}", th.handleDelete()).Methods("DELETE")
	mux.Handle("/{id:[0-9]+}", th.handlePut()).Methods("PUT")
}

// @Summary Get task
// @Security ApiKeyAuth
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task id"
// @Success	200 {object} model.Task
// @Failure 401 {object} errorResponce
// @Failure 404 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /tasks/{id} [get]
func (th *TaskHandler) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		task, err := th.store.Task().FindById(context.TODO(), int32(id))
		if err != nil {
			errorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		successRespond(w, r, http.StatusOK, task)
	}
}

// @Summary Create
// @Security ApiKeyAuth
// @Tags tasks
// @Accept json
// @Produce json
// @Param input body CreateInput true "task info"
// @Success	201	{object} model.Task
// @Failure 400 {object} errorResponce
// @Failure 401 {object} errorResponce
// @Failure 404 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /tasks [post]
func (th *TaskHandler) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &CreateInput{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		// todo-> соответсвует ли json структуре request?
		task := &model.Task{
			Owner: req.Owner,
			Name:  req.Name,
		}

		if err := task.Validate(); err != nil {
			errorRespond(w, r, http.StatusBadRequest, err)
			return
		}

		if err := th.store.Task().Create(context.TODO(), task); err != nil {
			errorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		successRespond(w, r, http.StatusCreated, task)
	}
}

// @Summary Delete task
// @Security ApiKeyAuth
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task id"
// @Success	200 {object} nil
// @Failure 401 {object} errorResponce
// @Failure 404 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /tasks/{id} [delete]
func (th *TaskHandler) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		if err := th.store.Task().DeleteById(context.TODO(), int32(id)); err != nil {
			errorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		successRespond(w, r, http.StatusOK, nil)
	}
}

// @Summary Update task
// @Security ApiKeyAuth
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task id"
// @Param input body UpdateInput true "task info"
// @Success	200 {object} model.Task
// @Failure 400 {object} errorResponce
// @Failure 401 {object} errorResponce
// @Failure 404 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /tasks/{id} [put]
func (th *TaskHandler) handlePut() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		req := &UpdateInput{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorRespond(w, r, http.StatusBadRequest, err)
			return
		}
		// todo-> соответсвует ли json структуре request?
		task := &model.Task{
			Owner:      req.Owner,
			Name:       req.Name,
			IsComplete: req.IsComplete,
		}

		if err := task.Validate(); err != nil {
			errorRespond(w, r, http.StatusBadRequest, err)
			return
		}

		if err := th.store.Task().UpdateById(context.TODO(), int32(id), task); err != nil {
			errorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		successRespond(w, r, http.StatusOK, task)
	}
}

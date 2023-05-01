package fake

import (
	"context"
	"errors"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
)

var errNoTaskWithThisId = errors.New("no task with this id")

type TaskRepository struct {
	tasks map[int32]*model.Task
}

var _ store.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[int32]*model.Task),
	}
}

func (tr *TaskRepository) Create(ctx context.Context, t *model.Task) error {
	t.ID = int32((len(tr.tasks) + 1))
	tr.tasks[t.ID] = t

	return nil
}

func (tr *TaskRepository) FindById(ctx context.Context, id int32) (*model.Task, error) {
	t, ok := tr.tasks[id]
	if !ok {
		return nil, errNoTaskWithThisId
	}

	return t, nil
}

func (tr *TaskRepository) DeleteById(ctx context.Context, id int32) error {
	_, ok := tr.tasks[id]
	if !ok {
		return errNoTaskWithThisId
	}

	delete(tr.tasks, id)
	return nil
}
func (tr *TaskRepository) UpdateById(ctx context.Context, id int32, t *model.Task) error {
	oldTask, ok := tr.tasks[id]
	if !ok {
		return errNoTaskWithThisId
	}

	t.ID = oldTask.ID
	tr.tasks[id] = t

	return nil
}

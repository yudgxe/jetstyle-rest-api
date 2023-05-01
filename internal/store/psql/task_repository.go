package psql

import (
	"context"
	"database/sql"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
)

type TaskRepository struct {
	db *sql.DB
}

var _ store.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) Create(ctx context.Context, t *model.Task) error {
	if err := tr.db.QueryRowContext(ctx,
		`INSERT INTO tasks(
			owner,
			name
		) VALUES ($1, $2) RETURNING id, create_date`,
		t.Owner,
		t.Name,
	).Scan(
		&t.ID,
		&t.CreateDate,
	); err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) FindById(ctx context.Context, id int32) (*model.Task, error) {
	task := &model.Task{}

	if err := tr.db.QueryRowContext(ctx,
		`SELECT id, create_date, update_date, owner, name, is_complete, complete_date FROM tasks WHERE id = $1`,
		id,
	).Scan(
		&task.ID,
		&task.CreateDate,
		&task.UpdateDate,
		&task.Owner,
		&task.Name,
		&task.IsComplete,
		&task.CompleteDate,
	); err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) DeleteById(ctx context.Context, id int32) error {
	var ID int32

	if err := tr.db.QueryRowContext(ctx,
		`DELETE FROM tasks WHERE id = $1 RETURNING id`,
		id,
	).Scan(&ID); err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) UpdateById(ctx context.Context, id int32, t *model.Task) error {
	if err := tr.db.QueryRowContext(ctx,
		`UPDATE tasks SET(owner, name, is_complete) = ($1, $2, $3) WHERE id = $4 RETURNING create_date, update_date, complete_date`,
		t.Owner,
		t.Name,
		t.IsComplete,
		id,
	).Scan(
		&t.CreateDate,
		&t.UpdateDate,
		&t.CompleteDate,
	); err != nil {
		return err
	}

	return nil
}

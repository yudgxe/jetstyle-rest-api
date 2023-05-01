package model

import (
	"errors"
	"time"
)

var errMinNameLenght = errors.New("insufficient name length. Minimum of 4 characters")

const minNameLenght = 4

type Task struct {
	ID           int32      `json:"id"    minimum:"1"`
	CreateDate   *time.Time `json:"create_date"`
	UpdateDate   *time.Time `json:"update_date"`
	Owner        int32      `json:"owner" minimum:"1"`
	Name         string     `json:"name"  minimum:"4"`
	IsComplete   bool       `json:"is_complete"`
	CompleteDate *time.Time `json:"complete_date"`
}

func (t *Task) Validate() error {
	if len(t.Name) < minNameLenght {
		return errMinNameLenght
	}

	return nil
}

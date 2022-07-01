package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type Task struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title" fako:"job_title" `
	Description string    `db:"description" fako:"sentence" `
	CreatedAt   time.Time `db:"created_at" `
	UpdatedAt   time.Time `db:"updated_at" `
	Must        time.Time `db:"must" `
	Status      bool      `db:"status" `
}

func (task *Task) IsValid(errors *validate.Errors) {
	if task.Title == "" {
		errors.Add("title", "Title must not be empty")
	}
	if task.Description == "" {
		errors.Add("description", "Description must not be empty")
	}
	if task.Must.String() == "" {
		errors.Add("must", "Must must not be empty")
	}

}

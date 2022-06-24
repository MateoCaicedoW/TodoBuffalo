package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
)

type Task struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Must        time.Time `db:"must"`
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

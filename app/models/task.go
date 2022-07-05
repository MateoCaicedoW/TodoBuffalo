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
	UserID      uuid.UUID `db:"user_id" `
	Must        time.Time `db:"must" `
	Status      bool      `db:"status" `
	User        *User     `belongs_to:"users"`
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
	if task.UserID == uuid.Nil {
		errors.Add("user_id", "User must not be empty")
	}

}

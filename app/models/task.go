package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
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

func (t *Task) Validate() (*validate.Errors, error) {

	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
		&validators.FuncValidator{
			Fn: func() bool {
				if t.UserID == uuid.Nil {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "UserID",
			Message: "%s User can't be blank.",
		},
		&validators.StringIsPresent{Field: t.Must.String(), Name: "Must"},
	), nil

}

package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type User struct {
	ID                   uuid.UUID `db:"id" `
	Email                string    `db:"email" fako:"email_address"`
	FirstName            string    `db:"first_name" fako:"first_name"`
	LastName             string    `db:"last_name" fako:"last_name"`
	Password             string    `db:"-" `
	PasswordHash         string    `db:"password_hash" fako:"password"`
	PasswordConfirmation string    `db:"-"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
	Task                 Task      `has_one:"tasks"`
}

func (user *User) IsValid(errors *validate.Errors) {

	if user.Email == "" {
		errors.Add("email", "Email must not be empty")
	}
	if user.Password == "" {
		errors.Add("password", "Password must not be empty")
	}
	if user.PasswordConfirmation == "" {
		errors.Add("password_confirmation", "Password confirmation must not be empty")
	}
	if user.Password != user.PasswordConfirmation {
		errors.Add("password_confirmation", "Password and Password Confirmation must match")
	}

	if user.FirstName == "" {
		errors.Add("first_name", "First Name must not be empty")
	}
	if user.LastName == "" {
		errors.Add("last_name", "Last Name must not be empty")
	}

}

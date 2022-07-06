package models

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
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

func (u *User) Validate(tx *pop.Connection, c buffalo.Context) (*validate.Errors, error) {

	return validate.Validate(
		&validators.StringIsPresent{Field: u.FirstName, Name: "First Name"},
		&validators.RegexMatch{Field: u.LastName, Name: "Last Name", Expr: `^[a-zA-Z]+$`, Message: "Last Name must be letters only"},
		&validators.RegexMatch{Field: u.FirstName, Name: "First Name", Expr: `^[a-zA-Z]+$`, Message: "First Name must be letters only"},
		&validators.StringIsPresent{Field: u.LastName, Name: "Last Name"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.FuncValidator{
			Fn: func() bool {
				if (c.Request().URL.String() == "/users/new/") && u.Password == "" {
					return false
				}
				return true
			},

			Name:    "Password",
			Message: "%s Password is required.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if (c.Request().URL.String() == "/users/new/" || c.Request().URL.String() != "/users/new/") && len(u.Password) > 0 {
					if u.Password != u.PasswordConfirmation {
						return false
					}
				}
				return true
			},

			Name:    "Password",
			Message: "%s Passwords do not match.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if (c.Request().URL.String() == "/users/new/") && u.Password != "" || (c.Request().URL.String() != "/users/new/" && u.Password != "") {
					if len(u.Password) < 8 {
						return false
					}
				}
				return true
			},
			Name:    "Password",
			Message: " %s Password must be at least 8 characters.",
		},

		&validators.FuncValidator{

			Name:    "Email",
			Message: "%s Email is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("email = ?", u.Email)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err := q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},

		&validators.EmailIsPresent{Name: "Email", Field: u.Email},
	), nil
}

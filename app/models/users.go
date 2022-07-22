package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	Rol = make(map[string]string)
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
	Rol                  string    `db:"rol"`
	Tasks                []Task    `has_many:"tasks"`
}

func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, *validate.Errors) {
	err := u.ValidatePass()
	err2, _ := u.Validate(tx)
	return err, err2
}
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, *validate.Errors) {
	err2, _ := u.Validate(tx)
	err := *validate.NewErrors()

	if u.Password != "" {
		err = *u.ValidatePass()
		return &err, err2
	}
	return &err, err2
}

func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Email = strings.ToLower(u.Email)
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), err
	}
	u.PasswordHash = string(ph)
	return tx.ValidateAndCreate(u)
}

func (u *User) ValidatePass() *validate.Errors {
	u.Password = strings.Replace(u.Password, " ", "", -1)
	u.PasswordConfirmation = strings.Replace(u.PasswordConfirmation, " ", "", -1)
	return validate.Validate(
		&validators.FuncValidator{
			Fn: func() bool {

				if len(u.Password) >= 8 && len(u.Password) <= 50 {
					if u.Password != "" && u.Password != u.PasswordConfirmation {
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
				if u.Password != "" {
					if len(u.Password) < 8 || len(u.Password) > 50 {
						return false
					}
				}
				return true
			},
			Name:    "Password",
			Message: " %s Password must be between 8 and 50 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Password == "" {
					return false
				}
				return true
			},

			Name:    "Password",
			Message: "%s Password is required.",
		},
	)
}

func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {

	return validate.Validate(
		&validators.StringIsPresent{Field: u.FirstName, Name: "First Name"},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.FirstName != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(u.FirstName) {
					return false
				}
				return true
			},
			Name:    "First Name",
			Message: "%s First Name must be letters only.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.LastName != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(u.LastName) {
					return false
				}
				return true
			},
			Name:    "Last Name",
			Message: "%s Last Name must be letters only.",
		},
		&validators.StringIsPresent{Field: u.LastName, Name: "Last Name"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},

		&validators.FuncValidator{

			Fn: func() bool {
				if u.FirstName != "" && len(u.FirstName) > 50 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FirstName) {
					return false
				}
				return true
			},
			Name:    "First Name",
			Message: "%s First Name must be less than 50 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.LastName != "" && len(u.LastName) > 50 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.LastName) {
					return false
				}
				return true
			},
			Name:    "Last Name",
			Message: "%s Last Name must be less than 50 characters.",
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
		&validators.FuncValidator{
			Fn: func() bool {
				ex := `^[\w\.]+@([\w-]+\.)+[\w-]{2,4}$`
				if u.Email != "" && !regexp.MustCompile(ex).MatchString(u.Email) {
					return false
				}
				return true
			},
			Name:    "Email",
			Message: "%s Email is invalid",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Email != "" {
					local := strings.Split(u.Email, "@")
					str := local[0]
					if len(str) > 64 {
						return false
					}
				}
				return true
			},
			Name:    "Email",
			Message: "%s Before @ Email must be less or equal than 64 characters ",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Email != "" {
					local := strings.Split(u.Email, "@")
					str := local[1]
					if len(str) > 255 {
						return false
					}
				}
				return true
			},
			Name:    "Email",
			Message: "%s After @ Email must be less or equal than 255 characters ",
		},
	), nil
}

package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID                   uuid.UUID `db:"id" `
	Email                string    `db:"email" fako:"email_address"`
	FirstName            string    `db:"first_name" fako:"first_name"`
	LastName             string    `db:"last_name" fako:"last_name"`
	Password             string    `db:"-" `
	PasswordHash         string    `db:"password_hash" fako:"password_hash"`
	PasswordConfirmation string    `db:"-"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}

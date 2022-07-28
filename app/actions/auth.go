package actions

import (
	"TodoBuffalo/app/models"
	"database/sql"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/markbates/errx"
	"golang.org/x/crypto/bcrypt"
)

// AuthNew loads the signin page
func AuthNew(c buffalo.Context) error {
	if c.Value("current_user") != nil {
		return c.Redirect(http.StatusSeeOther, "/todo")
	}

	c.Set("user", models.User{})
	return c.Render(http.StatusOK, r.HTML("auth/new.plush.html"))

}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {

	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	tx := c.Value("tx").(*pop.Connection)
	// find a user with the email
	err := tx.Where("email = ?", u.Email).First(u)

	if err != nil {
		if errx.Unwrap(err) == sql.ErrNoRows {

			// couldn't find an user with the supplied email address.
			return bad(c, u)
		}
		return err
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {

		return bad(c, u)
	}
	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome back to Todo!")

	return c.Redirect(http.StatusSeeOther, "/todo")
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func bad(c buffalo.Context, u *models.User) error {
	c.Set("user", u)
	verrs := validate.NewErrors()
	verrs.Add("email", "invalid email/password")
	c.Set("errors", verrs)
	return c.Render(http.StatusUnprocessableEntity, r.HTML("auth/new.plush.html"))
}

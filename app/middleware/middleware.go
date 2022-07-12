// middleware package is intended to host the middlewares used
// across the app.
package middleware

import (
	"TodoBuffalo/app/models"
	"database/sql"
	"net/http"

	"github.com/gobuffalo/buffalo"
	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/pop/v6"

	"github.com/markbates/errx"
	"github.com/wawandco/ox/pkg/buffalotools"
)

var (
	// RequestID pulls the request id from the request and
	// adds it if its not present.
	RequestID = buffalotools.NewRequestIDMiddleware("RequestID")

	// Database middleware adds a `tx` context variable
	// to every request, this tx variates to be a plain connection
	// or a transaction based on the type of request.
	Database = buffalotools.DatabaseMiddleware(models.DB(), nil)

	// ParameterLogger logs out parameters that the app received
	// taking care of sensitive data.
	ParameterLogger = paramlogger.ParameterLogger

	// CSRF middleware protects from CSRF attacks.
	CSRF = csrf.New
)

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				c.Session().Clear()
				if errx.Unwrap(err) == sql.ErrNoRows {
					return c.Redirect(http.StatusSeeOther, "/")
				}
				return err
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {

			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}

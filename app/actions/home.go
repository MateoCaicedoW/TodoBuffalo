package actions

import (
	"TodoBuffalo/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func HomeHandler(c buffalo.Context) error {
	if c.Value("current_user") != nil {
		return c.Redirect(http.StatusSeeOther, "/todo")
	}
	c.Set("user", models.User{})
	return c.Render(http.StatusOK, r.HTML("index.plush.html"))
}

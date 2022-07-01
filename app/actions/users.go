package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"

	"TodoBuffalo/app/models"
)

// UsersResource is the resource for the User model

// List gets all Users. This function is mapped to the path
// GET /users
func UsersList(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	users := []models.User{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	q = q.Order("created_at desc")
	// Retrieve all Users from the DB
	if err := q.All(&users); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("users", users)

	return c.Render(http.StatusOK, r.HTML("/users/index.plush.html"))
}

// // Show gets the data for one User. This function is mapped to
// // the path GET /users/{user_id}
// func UsersShow(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty User
// 	user := &models.User{}

// 	// To find the User the parameter user_id is used.
// 	if err := tx.Find(user, c.Param("user_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	c.Set("user", user)

// 	return c.Render(http.StatusOK, r.HTML("/users/show.plush.html"))
// }

// // New renders the form for creating a new User.
// // This function is mapped to the path GET /users/new
// func UsersNew(c buffalo.Context) error {
// 	c.Set("user", &models.User{})

// 	return c.Render(http.StatusOK, r.HTML("/users/new.plush.html"))
// }

// // Create adds a User to the DB. This function is mapped to the
// // path POST /users
func UsersCreate(c buffalo.Context) error {
	// Allocate an empty User
	user := &models.User{}

	// Bind user to the html form elements
	if err := c.Bind(user); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("user", user)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/users/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User created successfully")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/users")
}

// // Edit renders a edit form for a User. This function is
// // mapped to the path GET /users/{user_id}/edit
// func UsersEdit(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty User
// 	user := &models.User{}

// 	if err := tx.Find(user, c.Param("user_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	c.Set("user", user)

// 	return c.Render(http.StatusOK, r.HTML("/users/edit.plush.html"))
// }

// // Update changes a User in the DB. This function is mapped to
// // the path PUT /users/{user_id}
// func UsersUpdate(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty User
// 	user := &models.User{}

// 	if err := tx.Find(user, c.Param("user_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	// Bind User to the html form elements
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}

// 	verrs, err := tx.ValidateAndUpdate(user)
// 	if err != nil {
// 		return err
// 	}

// 	if verrs.HasAny() {
// 		// Make the errors available inside the html template
// 		c.Set("errors", verrs)

// 		// Render again the edit.html template that the user can
// 		// correct the input.
// 		c.Set("user", user)

// 		return c.Render(http.StatusUnprocessableEntity, r.HTML("/users/edit.plush.html"))
// 	}

// 	// If there are no errors set a success message
// 	c.Flash().Add("success", "user.updated.success")

// 	// and redirect to the show page
// 	return c.Redirect(http.StatusSeeOther, "userPath()", render.Data{"user_id": user.ID})
// }

// // Destroy deletes a User from the DB. This function is mapped
// // to the path DELETE /users/{user_id}
// func UsersDestroy(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty User
// 	user := &models.User{}

// 	// To find the User the parameter user_id is used.
// 	if err := tx.Find(user, c.Param("user_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	if err := tx.Destroy(user); err != nil {
// 		return err
// 	}

// 	// If there are no errors set a flash message
// 	c.Flash().Add("success", "user.destroyed.success")

// 	// Redirect to the index page
// 	return c.Redirect(http.StatusSeeOther, "usersPath()")
// }

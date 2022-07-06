package actions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"golang.org/x/crypto/bcrypt"

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

// // New renders the form for creating a new User.
// // This function is mapped to the path GET /users/new
func UsersShow(c buffalo.Context) error {

	// Allocate an empty User
	user := &models.User{}

	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("/users/new.plush.html"))
}

// // Create adds a User to the DB. This function is mapped to the
// // path POST /users
func UsersCreate(c buffalo.Context) error {
	// Allocate an empty User
	user := &models.User{}

	// Bind user to the html form elements
	if err := c.Bind(user); err != nil {
		return err
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the form input

	if verrs, _ := user.Validate(tx, c); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/users/new.plush.html"))
	}

	user.Email = strings.ToLower(user.Email)
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashPass)
	user.FirstName = strings.ToLower(user.FirstName)
	user.LastName = strings.ToLower(user.LastName)

	// Get the DB connection from the context

	// Validate the data from the html form
	err2 := tx.Create(user)
	if err2 != nil {
		return err2
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User created successfully")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/users")
}

// Edit renders a edit form for a User. This function is
// mapped to the path GET /users/{user_id}/edit
func UsersEdit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("/users/edit.plush.html"))
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func UsersUpdate(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind User to the html form elements
	if err := c.Bind(user); err != nil {
		return err
	}
	if verrs, _ := user.Validate(tx, c); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/users/edit.plush.html"))
	}

	userTemp := &models.User{}
	if err := tx.Find(userTemp, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Update(user); err != nil {
		return err
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User updated successfully")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/users")
}

// Destroy deletes a User from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func UsersDestroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "User deleted successfully")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "/users")
}

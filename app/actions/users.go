package actions

import (
	"fmt"
	"net/http"
	"strconv"
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
	keyword := "%" + strings.ToLower(c.Param("keyword")) + "%"
	if err := q.Where("lower(first_name) LIKE ? or lower(last_name) LIKE ? or lower(email) LIKE ? ", keyword, keyword, keyword).
		All(&users); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	fmt.Println(q.Paginator)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("users", users)
	setUsers(tx, c)
	SetTime(c)
	return c.Render(http.StatusOK, r.HTML("/users/index.plush.html"))
}

// // New renders the form for creating a new User.
// // This function is mapped to the path GET /users/new
func UsersNew(c buffalo.Context) error {

	a := c.Session().Get("current_user_id")

	if a != nil && c.Value("current_user").(*models.User).Rol != "admin" {
		c.Flash().Add("error", "You are not authorized to access this page")
		c.Redirect(http.StatusSeeOther, "/")
	}
	// Allocate an empty User
	user := &models.User{}

	setRol(c)
	SetTime(c)
	setUsers(c.Value("tx").(*pop.Connection), c)
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

	if verrs, verr2 := user.ValidateCreate(tx); verrs.HasAny() {
		verrs.Append(verr2)
		c.Set("errors", verrs)
		c.Set("user", user)
		setRol(c)
		SetTime(c)
		setUsers(tx, c)
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
	user.Rol = "user"

	// Validate the data from the html form
	err2 := tx.Eager().Create(user)
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
	setRol(c)
	setUsers(tx, c)
	SetTime(c)
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
	userTemp := &models.User{}
	if err := c.Bind(userTemp); err != nil {
		return err
	}
	userTemp.ID = user.ID

	if userTemp.Password == "" {
		userTemp.PasswordHash = user.PasswordHash
	}

	if userTemp.Password != "" {
		hashPass, err := bcrypt.GenerateFromPassword([]byte(userTemp.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userTemp.PasswordHash = string(hashPass)
	}

	verrs, verr2 := userTemp.ValidateUpdate(tx)
	verrs.Append(verr2)

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", userTemp)
		setRol(c)
		SetTime(c)
		setUsers(tx, c)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/users/edit.plush.html"))
	}

	if err := tx.Eager().Update(userTemp); err != nil {
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

	if err := tx.Eager().Destroy(user); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "User deleted successfully")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "/users")
}

func UsersShow(c buffalo.Context) error {
	user := &models.User{}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	if err := tx.Eager("Tasks").Find(user, c.Param("id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)
	count := len(user.Tasks)

	message := strconv.Itoa(count) + " Tasks"
	c.Set("message", message)
	c.Set("count", count)
	SetTime(c)
	return c.Render(http.StatusOK, r.HTML("/users/show.plush.html"))
}

func setRol(c buffalo.Context) {
	if c.Value("current_user") != nil {
		rol := models.Rol
		rol["user"] = "user"
		rol["admin"] = "admin"
		c.Set("rol", rol)
		c.Set("current_user", c.Value("current_user").(*models.User))
	}
}

func setUsers(tx *pop.Connection, c buffalo.Context) {
	users := models.User{}
	count, _ := tx.Count(users)

	message := strconv.Itoa(count) + " Users"
	c.Set("message", message)
}

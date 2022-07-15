package actions

import (
	"TodoBuffalo/app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

type TodoResource struct {
	buffalo.Resource
}

func (t TodoResource) List(c buffalo.Context) error {
	tasks := []models.Task{}

	tx := c.Value("tx").(*pop.Connection)

	q := tx.PaginateFromParams(c.Params())
	q = q.Order("created_at desc")

	u := c.Value("current_user").(*models.User)

	if u.Rol != "admin" {
		if err := q.Eager().Where("user_id = ?", u.ID).All(&tasks); err != nil {
			return err
		}
	}
	if u.Rol == "admin" {
		if err := q.Eager().All(&tasks); err != nil {
			return err
		}
	}
	c.Set("current_user", u)
	c.Set("tasks", tasks)
	c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.HTML("todo/index.plush.html"))
}

func (t TodoResource) New(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	var task models.Task
	task.Must = time.Now()

	findUsers(c, tx, users)

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("todo/new.plush.html"))
}

func (t TodoResource) Create(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	task := &models.Task{
		User: &models.User{},
	}

	if err := c.Bind(task); err != nil {

		return err
	}
	task.Status = false

	if err := validateCreateAndUpdate(c, task, tx, users); err != "" {
		return c.Render(http.StatusUnprocessableEntity, r.HTML("todo/new.plush.html"))
	}

	if err := tx.Eager().Create(task); err != nil {
		return err
	}

	c.Flash().Add("success", "Record was successfully created!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func (t TodoResource) Show(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	var task models.Task
	id := c.Param("todo_id")

	task.ID = uuid.FromStringOrNil(id)
	if err := tx.Eager().Find(&task, id); err != nil {
		return err
	}

	findUsers(c, tx, users)

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo/edit.plush.html"))
}

func (t TodoResource) Update(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	taskTemp := &models.Task{}

	id := c.Param("todo_id")

	taskTemp.ID = uuid.FromStringOrNil(id)
	if err := c.Bind(taskTemp); err != nil {

		return err
	}

	if err := validateCreateAndUpdate(c, taskTemp, tx, users); err != "" {

		return c.Render(http.StatusUnprocessableEntity, r.HTML("todo/edit.plush.html"))
	}

	if err := tx.Eager().Update(taskTemp); err != nil {
		return err
	}

	fmt.Println("Pero aqui si")
	c.Flash().Add("success", "Record was successfully updated!")
	return c.Redirect(http.StatusSeeOther, "/todo")
}

func (t TodoResource) Destroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	taskTemp := &models.Task{}
	id := c.Param("todo_id")
	taskTemp.ID = uuid.FromStringOrNil(id)

	if err := tx.Eager().Destroy(taskTemp); err != nil {
		return err
	}
	c.Flash().Add("success", "Record was successfully deleted!")
	return c.Redirect(http.StatusSeeOther, "/todo")
}

func Status(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	taskTemp := &models.Task{}
	id := c.Param("todo_id")
	taskTemp.ID = uuid.FromStringOrNil(id)
	if err := tx.Eager().Find(taskTemp, id); err != nil {
		return err
	}
	taskTemp.Status = !taskTemp.Status
	if err := tx.Eager().Update(taskTemp); err != nil {
		return err
	}
	if taskTemp.Status {
		c.Flash().Add("success", "Record was successfully completed!")
	}

	if !taskTemp.Status {
		c.Flash().Add("success", "Record was successfully uncompleted!")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func findUsers(c buffalo.Context, tx *pop.Connection, users []models.User) {

	//create map of users
	userMap := make(map[string]interface{})

	for i := 0; i < len(users); i++ {

		userMap[users[i].Email] = users[i].ID.String()
	}
	c.Set("users", userMap)

}

func validateCreateAndUpdate(c buffalo.Context, task *models.Task, tx *pop.Connection, arrayUsers []models.User) string {
	user := c.Value("current_user").(*models.User)
	if user.Rol != "admin" {
		task.User = user
		task.UserID = user.ID
	}

	if verrs, _ := task.Validate(); verrs.HasAny() {
		findUsers(c, tx, arrayUsers)

		c.Set("task", task)
		c.Set("errors", verrs.Errors)

		return "error"
	}
	return ""
}

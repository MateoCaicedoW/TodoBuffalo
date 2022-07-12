package actions

import (
	"TodoBuffalo/app/models"
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
	if err := q.Eager().Where("user_id = ?", u.ID).All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)
	c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.HTML("todo/index.plush.html"))
}

func (t TodoResource) New(c buffalo.Context) error {
	users := []models.User{}
	tx := c.Value("tx").(*pop.Connection)
	var task models.Task
	task.Must = time.Now()
	if err := tx.All(&users); err != nil {
		return err
	}

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("todo/new.plush.html"))
}

func (t TodoResource) Create(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	task := &models.Task{
		User: &models.User{},
	}

	if err := c.Bind(task); err != nil {

		return err
	}
	task.Status = false
	user := c.Value("current_user").(*models.User)
	task.User = user
	task.UserID = user.ID

	if verrs, _ := task.Validate(); verrs.HasAny() {
		c.Set("task", task)
		users := []models.User{}
		if err := tx.All(&users); err != nil {
			return err
		}
		// create map of users

		c.Set("errors", verrs.Errors)

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
	var task models.Task
	id := c.Param("todo_id")

	task.ID = uuid.FromStringOrNil(id)
	if err := tx.Eager().Find(&task, id); err != nil {
		return err
	}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo/edit.plush.html"))
}

func (t TodoResource) Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	taskTemp := &models.Task{}
	id := c.Param("todo_id")
	taskTemp.ID = uuid.FromStringOrNil(id)
	if err := c.Bind(taskTemp); err != nil {
		return err
	}
	user := c.Value("current_user").(*models.User)
	taskTemp.User = user
	taskTemp.UserID = user.ID

	if verrs, _ := taskTemp.Validate(); verrs.HasAny() {
		c.Set("task", taskTemp)
		c.Set("errors", verrs.Errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("todo/edit.plush.html"))
	}
	if err := tx.Eager().Update(taskTemp); err != nil {
		return err
	}

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

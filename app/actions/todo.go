package actions

import (
	"TodoBuffalo/app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

func Index(c buffalo.Context) error {
	tasks := []models.Task{}
	tx := c.Value("tx").(*pop.Connection)

	q := tx.PaginateFromParams(c.Params())
	q = q.Order("created_at desc")

	if err := q.Eager().All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)
	c.Set("pagination", q.Paginator)

	url := c.Request().URL.String()
	fmt.Println(url)
	c.Set("url", url)
	return c.Render(http.StatusOK, r.HTML("todo/index.plush.html"))
}

func New(c buffalo.Context) error {
	users := []models.User{}
	tx := c.Value("tx").(*pop.Connection)
	var task models.Task
	task.Must = time.Now()
	if err := tx.All(&users); err != nil {
		return err
	}

	c.Set("users", users)
	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo/new.plush.html"))
}

func Create(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	task := &models.Task{
		User: &models.User{},
	}

	if err := c.Bind(task); err != nil {

		return err
	}
	task.Status = false

	err := validate.Validate(task)

	for item := range err.Errors {
		c.Flash().Add("error", err.Errors[item][0])
		c.Set("task", task)
		users := []models.User{}
		if err := tx.All(&users); err != nil {
			return err
		}

		c.Set("users", users)
		return c.Render(http.StatusBadRequest, r.HTML("todo/new.plush.html"))
	}
	if err := tx.Eager().Create(task); err != nil {
		return err
	}

	c.Flash().Add("success", "Record was successfully created!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func Edit(c buffalo.Context) error {
	users := []models.User{}
	tx := c.Value("tx").(*pop.Connection)
	var task models.Task
	id := c.Param("id")
	task.ID = uuid.FromStringOrNil(id)
	if err := tx.Eager().Find(&task, id); err != nil {
		return err
	}
	if err := tx.All(&users); err != nil {
		return err
	}
	c.Set("users", users)
	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo/edit.plush.html"))
}

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	taskTemp := &models.Task{}
	id := c.Param("id")
	taskTemp.ID = uuid.FromStringOrNil(id)
	if err := c.Bind(taskTemp); err != nil {
		return err
	}

	err := validate.Validate(taskTemp)

	for item := range err.Errors {
		c.Flash().Add("error", err.Errors[item][0])
		c.Set("task", taskTemp)
		return c.Render(http.StatusBadRequest, r.HTML("todo/edit.plush.html"))
	}

	if err := tx.Eager().Update(taskTemp); err != nil {
		return err
	}

	c.Flash().Add("success", "Record was successfully updated!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	taskTemp := &models.Task{}
	id := c.Param("id")
	taskTemp.ID = uuid.FromStringOrNil(id)

	if err := tx.Eager().Destroy(taskTemp); err != nil {
		return err
	}
	c.Flash().Add("success", "Record was successfully deleted!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func Status(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	taskTemp := &models.Task{}
	id := c.Param("id")
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

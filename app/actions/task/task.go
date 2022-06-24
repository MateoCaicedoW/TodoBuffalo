package task

import (
	"TodoBuffalo/app/models"
	"TodoBuffalo/app/render"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func Index(c buffalo.Context) error {
	tasks := []models.Task{}
	if err := models.DB().All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)

	return c.Render(http.StatusOK, r.HTML("tasks/index.plush.html"))
}

func New(c buffalo.Context) error {
	var task models.Task
	task.Must = time.Now()
	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("tasks/new.plush.html"))
}
func Create(c buffalo.Context) error {
	task := &models.Task{}
	task.ID = uuid.Must(uuid.NewV4()).String()
	if err := c.Bind(task); err != nil {

		return err
	}

	err := validate.Validate(task)

	for item := range err.Errors {
		c.Flash().Add("error", err.Errors[item][0])
		c.Set("task", task)
		return c.Render(http.StatusBadRequest, r.HTML("tasks/new.plush.html"))
	}
	if err := models.DB().Create(task); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func Edit(c buffalo.Context) error {
	var task models.Task
	id := c.Param("id")
	task.ID = id
	if err := models.DB().Find(&task, id); err != nil {
		return err
	}

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("tasks/edit.plush.html"))
}

func Update(c buffalo.Context) error {
	taskTemp := &models.Task{}
	id := c.Param("id")
	taskTemp.ID = id

	if err := c.Bind(taskTemp); err != nil {
		return err
	}

	err := validate.Validate(taskTemp)

	for item := range err.Errors {
		c.Flash().Add("error", err.Errors[item][0])
		c.Set("task", taskTemp)
		return c.Render(http.StatusBadRequest, r.HTML("tasks/edit.plush.html"))
	}

	if err := models.DB().Update(taskTemp); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c buffalo.Context) error {
	taskTemp := &models.Task{}
	id := c.Param("id")
	taskTemp.ID = id

	if err := models.DB().Destroy(taskTemp); err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

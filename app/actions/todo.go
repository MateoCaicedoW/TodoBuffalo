package actions

import (
	"TodoBuffalo/app/models"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

func Index(c buffalo.Context) error {
	tasks := []models.Task{}

	// if err := models.DB().All(&tasks); err != nil {
	// 	return err
	// }

	// q := models.DB().Paginate(1, 10).Order("created_at desc")
	// q.All(&tasks)
	// fmt.Println(q.Paginator.TotalPages)

	q := models.DB().PaginateFromParams(c.Params())
	q = q.Order("created_at desc")

	if err := q.All(&tasks); err != nil {
		return err
	}
	c.Set("tasks", tasks)
	c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.HTML("todo/index.plush.html"))
}

func New(c buffalo.Context) error {
	var task models.Task
	task.Must = time.Now()
	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo/new.plush.html"))
}
func Create(c buffalo.Context) error {
	task := &models.Task{}
	if err := c.Bind(task); err != nil {
		return err
	}
	err := validate.Validate(task)
	for item := range err.Errors {
		c.Flash().Add("error", err.Errors[item][0])
		c.Set("task", task)
		return c.Render(http.StatusBadRequest, r.HTML("todo/new.plush.html"))
	}
	if err := models.DB().Create(task); err != nil {

		return err
	}
	c.Flash().Add("success", "Record was successfully created!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func Edit(c buffalo.Context) error {
	var task models.Task
	id := c.Param("id")
	task.ID = uuid.FromStringOrNil(id)
	if err := models.DB().Find(&task, id); err != nil {
		return err
	}

	c.Set("task", task)

	return c.Render(http.StatusOK, r.HTML("todo/edit.plush.html"))
}

func Update(c buffalo.Context) error {
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

	if err := models.DB().Update(taskTemp); err != nil {
		return err
	}

	c.Flash().Add("success", "Record was successfully updated!")
	return c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c buffalo.Context) error {
	taskTemp := &models.Task{}
	id := c.Param("id")
	taskTemp.ID = uuid.FromStringOrNil(id)

	if err := models.DB().Destroy(taskTemp); err != nil {
		return err
	}
	c.Flash().Add("success", "Record was successfully deleted!")
	return c.Redirect(http.StatusSeeOther, "/")
}

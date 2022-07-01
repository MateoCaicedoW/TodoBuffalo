package tasks

import (
	"TodoBuffalo/app/models"
	"time"

	"github.com/markbates/grift/grift"
	"github.com/wawandco/fako"
)

var _ = grift.Add("create:task", func(c *grift.Context) error {

	for i := 0; i < 100; i++ {
		var task models.Task
		fako.Fill(&task)
		task.Must = time.Now()
		task.Status = false
		if err := models.DB().Create(&task); err != nil {
			return err
		}
	}
	return nil
})

var _ = grift.Add("create:users", func(c *grift.Context) error {

	for i := 0; i < 30; i++ {
		var user models.User
		fako.Fill(&user)
		if err := models.DB().Create(&user); err != nil {
			return err
		}
	}
	return nil
})

package tasks

import (
	"TodoBuffalo/app/models"
	"time"

	"github.com/gofrs/uuid"
	"github.com/markbates/grift/grift"
	"github.com/wawandco/fako"
)

var _ = grift.Add("create:task", func(c *grift.Context) error {

	for i := 0; i < 100; i++ {
		var task models.Task
		fako.FillExcept(&task, "ID")
		task.ID = uuid.Must(uuid.NewV4()).String()
		task.Must = time.Now()

		if err := models.DB().Create(&task); err != nil {
			return err
		}
	}
	return nil
})

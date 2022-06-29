package tasks

import (
	"TodoBuffalo/app/models"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/markbates/grift/grift"
)

var _ = grift.Add("create:task", func(c *grift.Context) error {
	task := &models.Task{
		ID:          uuid.Must(uuid.NewV4()).String(),
		Title:       "Todo",
		Description: "Todo description",
		Must:        time.Now(),
	}

	for i := 0; i < 100; i++ {
		task.Title = task.Title + strconv.Itoa(i)
		task.Description = task.Description + strconv.Itoa(i)
		models.DB().Create(task)
	}
	return nil
})

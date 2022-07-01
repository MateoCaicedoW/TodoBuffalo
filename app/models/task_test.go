package models

import (
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/wawandco/fako"
)

func (ms *ModelSuite) Test_Create_Error() {
	task := &Task{
		Title: "Test Task",
		Must:  time.Now(),
	}

	err := validate.Validate(task)
	ms.Error(err)
	ms.True(err.HasAny())

	count, errs := ms.DB.Count("tasks")

	ms.NoError(errs)
	ms.Equal(0, count)

}

func (ms *ModelSuite) Test_Create_Success() {
	task := &Task{}
	fako.Fill(task)
	vers, err := ms.DB.ValidateAndCreate(task)
	ms.NoError(err)
	ms.False(vers.HasAny())

}

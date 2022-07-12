package models

import (
	"time"

	"github.com/wawandco/fako"
)

func (ms *ModelSuite) Test_Create_Error() {

	task := &Task{
		Title: "Test Task",
		Must:  time.Now(),
	}

	err, _ := task.Validate()
	ms.Error(err)
	ms.True(err.HasAny())

	count, errs := ms.DB.Count("tasks")

	ms.NoError(errs)
	ms.Equal(0, count)

}

func (ms *ModelSuite) Test_Create_Success() {
	user := &User{}
	fako.Fill(user)
	err1 := ms.DB.Create(user)
	ms.NoError(err1)

	task := &Task{}
	fako.Fill(task)
	task.UserID = user.ID
	vers, err := ms.DB.ValidateAndCreate(task)

	ms.NoError(err)
	ms.False(vers.HasAny())

}

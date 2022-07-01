package actions_test

import (
	"TodoBuffalo/app/models"
	"time"

	"github.com/gofrs/uuid"
	"github.com/wawandco/fako"
)

func (as *ActionSuite) Test_Index() {
	tasks := [2]models.Task{}

	for i := 0; i < len(tasks); i++ {
		fako.Fill(&tasks[i])
		tasks[i].Must = time.Now()
		err := as.DB.Create(&tasks[i])
		as.NoError(err)
	}

	res := as.HTML("/").Get()
	body := res.Body.String()
	for _, t := range tasks {
		as.Contains(body, t.Title)
	}

}

func (as *ActionSuite) Test_New() {
	res := as.HTML("/new").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, "New Task")
}

func (as *ActionSuite) Test_Create() {
	task := &models.Task{
		ID:          uuid.Must(uuid.NewV4()),
		Title:       "Test Task",
		Description: "Test Description",
		Must:        time.Now(),
	}

	res := as.HTML("/new/").Post(task)
	as.Equal(303, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Edit() {
	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	err := as.DB.Create(task)
	as.NoError(err)

	res := as.HTML("/edit/" + task.ID.String()).Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, task.Title)
	as.Contains(body, "Edit Task")
}

func (as *ActionSuite) Test_Update() {
	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	err := as.DB.Create(task)
	as.NoError(err)

	taskUpdate := &models.Task{}
	fako.Fill(taskUpdate)
	taskUpdate.ID = task.ID

	res := as.HTML("/edit/" + task.ID.String()).Put(taskUpdate)
	as.Equal(303, res.Code)
	as.Equal("/", res.Location())
	as.DB.Reload(task)
	as.Equal(taskUpdate.Title, task.Title)
}

func (as *ActionSuite) Test_Destroy() {
	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	err := as.DB.Create(task)
	as.NoError(err)

	res := as.HTML("/delete/" + task.ID.String()).Delete()
	as.Equal(303, res.Code)
	as.Equal("/", res.Location())

}

func (as *ActionSuite) Test_Status() {
	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	err := as.DB.Create(task)
	as.NoError(err)

	taskUpdate := &models.Task{}
	fako.Fill(taskUpdate)
	taskUpdate.ID = task.ID
	taskUpdate.Status = true

	res := as.HTML("/status/" + task.ID.String()).Put(taskUpdate)
	as.Equal(303, res.Code)
	as.Equal("/", res.Location())
}

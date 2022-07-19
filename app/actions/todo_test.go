package actions_test

import (
	"TodoBuffalo/app/models"
	"net/http"

	"time"

	"github.com/gofrs/uuid"
	"github.com/wawandco/fako"
)

func (as *ActionSuite) Test_Index() {
	tasks := [2]models.Task{}
	users := [2]models.User{}

	setUSerAdmin(as)
	for i := 0; i < len(tasks); i++ {
		fako.Fill(&tasks[i])
		fako.Fill(&users[i])
		err1 := as.DB.Create(&users[i])
		as.NoError(err1)
		tasks[i].UserID = users[i].ID
		tasks[i].Must = time.Now()
		err := as.DB.Create(&tasks[i])
		as.NoError(err)
	}

	res := as.HTML("/todo").Get()
	body := res.Body.String()

	for _, t := range tasks {
		as.Contains(body, t.Title)
	}

}

func (as *ActionSuite) Test_Failed_Index() {
	res := as.HTML("/todo").Get()
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_New() {
	setUser(as)
	res := as.HTML("/todo/new").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, "New Task")
}

func (as *ActionSuite) Test_Create() {

	setUser(as)
	user := &models.User{}
	fako.Fill(user)
	err := as.DB.Create(user)
	as.NoError(err)
	task := &models.Task{
		ID:          uuid.Must(uuid.NewV4()),
		Title:       "Test Task",
		Description: "Test Description",
		Must:        time.Now(),
		UserID:      user.ID,
		User:        user,
	}

	res := as.HTML("/todo/").Post(task)
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Edit() {
	setUser(as)
	user := &models.User{}
	fako.Fill(user)
	err1 := as.DB.Create(user)
	as.NoError(err1)

	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	task.UserID = user.ID
	err := as.DB.Create(task)
	as.NoError(err)

	res := as.HTML("/todo/" + task.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, task.Title)
	as.Contains(body, "Edit Task")
}

func (as *ActionSuite) Test_Update() {
	setUser(as)
	user := &models.User{}
	fako.Fill(user)
	err1 := as.DB.Create(user)
	as.NoError(err1)

	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	task.UserID = user.ID
	err := as.DB.Create(task)
	as.NoError(err)

	taskUpdate := &models.Task{}
	fako.Fill(taskUpdate)
	taskUpdate.UserID = user.ID
	taskUpdate.ID = task.ID
	taskUpdate.Must = time.Now()

	res := as.HTML("/todo/" + task.ID.String()).Put(taskUpdate)

	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/todo", res.Location())
	as.DB.Reload(task)
	as.Equal(taskUpdate.Title, task.Title)
}

func (as *ActionSuite) Test_Destroy() {
	setUser(as)
	user := &models.User{}
	fako.Fill(user)
	err1 := as.DB.Create(user)
	as.NoError(err1)

	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	task.UserID = user.ID
	err := as.DB.Create(task)
	as.NoError(err)

	res := as.HTML("/todo/" + task.ID.String()).Delete()
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/todo", res.Location())

}

func (as *ActionSuite) Test_Status() {
	setUser(as)
	user := &models.User{}
	fako.Fill(user)
	err1 := as.DB.Create(user)
	as.NoError(err1)

	task := &models.Task{}
	fako.Fill(task)
	task.Must = time.Now()
	task.UserID = user.ID
	err := as.DB.Create(task)
	as.NoError(err)

	taskUpdate := &models.Task{}
	fako.Fill(taskUpdate)
	taskUpdate.ID = task.ID
	taskUpdate.Status = true

	res := as.HTML("/todo/status/" + task.ID.String()).Put(taskUpdate)
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/", res.Location())
}

func setUser(as *ActionSuite) {
	u := &models.User{
		FirstName:            "John",
		LastName:             "Doe",
		Email:                "caicedomateo9@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Rol:                  "user",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_user_id", u.ID)
	as.Session.Set("current_user", u)
}

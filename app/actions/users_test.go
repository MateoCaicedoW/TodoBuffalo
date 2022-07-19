package actions_test

import (
	"TodoBuffalo/app/models"
	"net/http"

	"github.com/wawandco/fako"
)

func (as *ActionSuite) Test_ListUsers() {
	setUSerAdmin(as)
	users := [2]models.User{}
	for i := 0; i < len(users); i++ {
		fako.Fill(&users[i])
		err := as.DB.Create(&users[i])
		as.NoError(err)
	}
	res := as.HTML("/users").Get()
	body := res.Body.String()
	for _, u := range users {
		as.Contains(body, u.Email)
	}

}
func (as *ActionSuite) Test_ListUsers_Failed() {
	setUser(as)
	res := as.HTML("/users").Get()
	as.Equal(http.StatusSeeOther, res.Code)

}

func (as *ActionSuite) Test_NewUser() {
	setUSerAdmin(as)
	res := as.HTML("/users/new").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, "New User")
}

func (as *ActionSuite) Test_CreateUser() {
	setUSerAdmin(as)
	u := &models.User{
		FirstName:            "Test",
		LastName:             "User",
		Email:                "test@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	res := as.HTML("/users/new").Post(u)
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/users", res.Location())
}
func (as *ActionSuite) Test_NewUser_Failed() {
	setUser(as)
	res := as.HTML("/users/new").Get()
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_EditUser() {
	setUSerAdmin(as)
	u := &models.User{
		FirstName:            "Test",
		LastName:             "User",
		Email:                "test@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	err := as.DB.Create(u)
	as.NoError(err)
	res := as.HTML("/users/edit/" + u.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, "Edit User")

}

func (as *ActionSuite) Test_UpdateUSer() {
	setUSerAdmin(as)
	u := &models.User{
		FirstName:            "Test",
		LastName:             "User",
		Email:                "test@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Rol:                  "user",
	}
	err := as.DB.Create(u)
	as.NoError(err)

	usersUpdate := u
	usersUpdate.FirstName = "TestUpdate"

	res := as.HTML("/users/edit/" + u.ID.String()).Put(usersUpdate)

	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/users", res.Location())
	as.DB.Reload(u)
	as.Equal(u.FirstName, usersUpdate.FirstName)

}

func (as *ActionSuite) Test_DestroyUser() {
	setUSerAdmin(as)
	u := &models.User{
		FirstName:            "Test",
		LastName:             "User",
		Email:                "test@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	err := as.DB.Create(u)
	as.NoError(err)

	res := as.HTML("/users/delete/" + u.ID.String()).Delete()
	as.Equal(http.StatusSeeOther, res.Code)

}

func (as *ActionSuite) Test_SHowUser() {
	setUSerAdmin(as)
	u := &models.User{
		FirstName:            "Test",
		LastName:             "User",
		Email:                "test@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	err := as.DB.Create(u)
	as.NoError(err)
	res := as.HTML("/users/show/" + u.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Tasks")

}

func setUSerAdmin(as *ActionSuite) {
	u := &models.User{
		FirstName:            "John",
		LastName:             "Doe",
		Email:                "admin@gmail.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Rol:                  "admin",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_user_id", u.ID)
	as.Session.Set("current_user", u)
}

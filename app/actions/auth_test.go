package actions_test

import (
	"TodoBuffalo/app/models"
	"net/http"
)

func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/signin").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, "Log In")
}

func (as *ActionSuite) Test_Auth_Create() {

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

	res := as.HTML("/signin").Post(u)
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/todo", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
	u := &models.User{
		Email:    "mark@example.com",
		Password: "password",
	}
	res := as.HTML("/signin").Post(u)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
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

	u.Password = "passwordbad"
	res := as.HTML("/signin").Post(u)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

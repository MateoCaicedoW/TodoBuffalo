package actions_test

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/signin", res.Location())
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	setUser(as)
	res := as.HTML("/").Get()
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/todo", res.Location())

	as.Session.Clear()
	res = as.HTML("/").Get()
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/signin", res.Location())
}

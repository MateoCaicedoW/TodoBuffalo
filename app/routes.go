package app

import (
	"TodoBuffalo/app/actions"
	"TodoBuffalo/app/middleware"
	"TodoBuffalo/public"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", actions.HomeHandler)

	root.Use(actions.SetCurrentUser)
	root.Use(actions.Authorize)

	root.GET("/signin", actions.AuthNew)
	root.POST("/signin", actions.AuthCreate)
	root.DELETE("/signout", actions.AuthDestroy)
	root.GET("/users", actions.UsersList)
	root.GET("/users/new", actions.UsersShow)
	root.POST("/users/new", actions.UsersCreate)
	root.GET("/users/edit/{id}", actions.UsersEdit)
	root.PUT("/users/edit/{id}", actions.UsersUpdate)
	root.DELETE("/users/delete/{id}", actions.UsersDestroy)

	root.Middleware.Skip(actions.Authorize, actions.HomeHandler, actions.UsersShow, actions.UsersCreate, actions.AuthNew, actions.AuthCreate)
	root.Resource("/todo", actions.TodoResource{})
	root.PUT("/todo/status/{todo_id}", actions.Status)
	root.ServeFiles("/", http.FS(public.FS()))
}

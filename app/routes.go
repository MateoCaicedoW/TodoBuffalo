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

	root.Use(middleware.Authorize)
	root.Use(middleware.SetCurrentUser)

	root.GET("/signin", actions.AuthNew)
	root.POST("/signin", actions.AuthCreate)
	root.DELETE("/signout", actions.AuthDestroy)
	root.GET("/users/new", actions.UsersNew)
	root.POST("/users/new", actions.UsersCreate)
	root.GET("/users", middleware.MyMiddleware(actions.UsersList))
	root.GET("/users/edit/{id}", middleware.MyMiddleware(actions.UsersEdit))
	root.PUT("/users/edit/{id}", middleware.MyMiddleware(actions.UsersUpdate))
	root.GET("/users/show/{id}", middleware.MyMiddleware(actions.UsersShow))
	root.DELETE("/users/delete/{id}", middleware.MyMiddleware(actions.UsersDestroy))

	root.Middleware.Skip(middleware.Authorize, actions.HomeHandler, actions.AuthNew, actions.AuthCreate, actions.UsersNew, actions.UsersCreate)

	root.Resource("/todo", actions.TodoResource{})
	root.PUT("/todo/status/{todo_id}", actions.Status)
	root.GET("/todo/show/{todo_id}", actions.ShowInformation)
	root.ServeFiles("/", http.FS(public.FS()))
}

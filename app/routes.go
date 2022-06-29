package app

import (
	"TodoBuffalo/app/actions/todo"
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

	root.GET("/", todo.Index)
	root.GET("/new", todo.New)
	root.GET("/edit/{id}", todo.Edit)
	root.POST("/new", todo.Create)
	root.PUT("/edit/{id}", todo.Update)
	root.DELETE("/delete/{id}", todo.Delete)
	root.ServeFiles("/", http.FS(public.FS()))
}

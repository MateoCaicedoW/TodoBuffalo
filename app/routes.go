package app

import (
	"TodoBuffalo/app/actions/task"
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

	root.GET("/", task.Index)
	root.GET("/new", task.New)
	root.GET("/edit/{id}", task.Edit)
	root.POST("/new", task.Create)
	root.PUT("/edit/{id}", task.Update)
	root.DELETE("/delete/{id}", task.Delete)
	root.ServeFiles("/", http.FS(public.FS()))
}

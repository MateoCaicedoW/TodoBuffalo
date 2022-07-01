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

	root.GET("/", actions.Index)
	root.GET("/new", actions.New)
	root.GET("/edit/{id}", actions.Edit)
	root.POST("/new", actions.Create)
	root.PUT("/edit/{id}", actions.Update)
	root.DELETE("/delete/{id}", actions.Delete)
	root.ServeFiles("/", http.FS(public.FS()))
}

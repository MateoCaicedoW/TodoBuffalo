package render

import (
	"TodoBuffalo/app/models"
	"TodoBuffalo/app/templates"
	"TodoBuffalo/public"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/plush/v4"
	"github.com/wawandco/ox/pkg/buffalotools"
)

// Engine for rendering across the app, it provides
// the base for rendering HTML, JSON, XML and other formats
// while also defining thing like the base layout.
var Engine = render.New(render.Options{
	HTMLLayout:  "application.plush.html",
	TemplatesFS: templates.FS(),
	AssetsFS:    public.FS(),
	Helpers:     Helpers,
})

// Helpers available for the plush templates, there are
// some helpers that are injected by Buffalo but this is
// the list of custom Helpers.
var Helpers = map[string]interface{}{
	// partialFeeder is the helper used by the render engine
	// to find the partials that will be used, this is important
	"partialFeeder": buffalotools.NewPartialFeeder(templates.FS()),
	"currentUser":   currentUser,
	"admin":         Admin,
	"adminPermiss":  AdminPermiss,
	"redirect":      Redirect,
	"isEdit": func(i models.Task) string {
		if i.Status == true {
			return "d-none"
		}
		return "d-block"

	},
	"isActive": activeClass,
}

func currentUser(c plush.HelperContext) bool {

	if c.Value("current_user") != nil {
		return true
	}

	return false
}

func Admin(c plush.HelperContext) bool {

	if c.Value("current_user") != nil {
		user := c.Value("current_user").(*models.User)
		if user.Rol == "admin" {
			return true
		}
	}

	return false
}

func AdminPermiss(c plush.HelperContext) string {
	if Admin(c) {
		return "col-md-6 pe-0 pe-md-2"
	}
	return "col-md-12 pe-0"
}

func Redirect(c plush.HelperContext) string {
	if Admin(c) {
		return "/users"
	}
	return "/"
}

func activeClass(n string, help plush.HelperContext) string {

	if p, ok := help.Value("current_route").(buffalo.RouteInfo); ok {

		if p.PathName == n {
			return "custom-active"
		}
	}
	return ""
}

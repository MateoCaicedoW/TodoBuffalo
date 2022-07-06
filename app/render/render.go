package render

import (
	"TodoBuffalo/app/templates"
	"TodoBuffalo/public"

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
	"isActive":      isActive,
}

func isActive(name string, c plush.HelperContext) string {

	if cr, ok := c.Value("current_path").(string); ok {

		if cr == name {
			return "custom-active"
		}
	}

	return ""
}

package actions

import "github.com/gobuffalo/buffalo"

// ExtraHandler is a default handler to serve up
// a home page.
func ExtraHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("extra.html"))
}

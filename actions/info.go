package actions

import "github.com/gobuffalo/buffalo"

// InfoHandler is a default handler to serve up
// a home page.
func InfoHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("info.html"))
}

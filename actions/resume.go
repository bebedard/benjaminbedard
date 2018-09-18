package actions

import "github.com/gobuffalo/buffalo"

// ResumeHandler is a default handler to serve up
// a home page.
func ResumeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("resume.html"))
}

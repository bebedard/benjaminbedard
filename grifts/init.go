package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/benjaminbedard/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}

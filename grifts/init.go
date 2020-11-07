package grifts

import (
	"github.com/bhaveshitankar/my_show/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}

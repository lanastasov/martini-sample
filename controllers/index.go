package controllers

import (
	"github.com/martini-contrib/render"
)

func ShowIndex(r render.Render) {
	r.HTML(200, "index", struct {
		Message string
	}{
		Message: "index",
	})
}

package main

import (
	"github.com/go-martini/martini"
	"github.com/m0t0k1ch1/martini-sample/controllers"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", controllers.ShowIndex)

	m.Run()
}

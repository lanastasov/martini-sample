package main

import (
	"github.com/go-martini/martini"
	"github.com/m0t0k1ch1/martini-sample/controllers"
	"github.com/m0t0k1ch1/martini-sample/models"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	m.Map(db)

	m.Get("/", controllers.ShowIndex)

	m.Run()
}

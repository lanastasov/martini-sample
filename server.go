package main

import (
	"github.com/coopernurse/gorp"
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

	db, err := models.InitDB("sqlite3", "./data/martini-sample.db", gorp.SqliteDialect{})
	if err != nil {
		panic(err)
	}
	m.Map(db)

	m.Get("/", controllers.ShowIndex)
	m.Get("/thread/list", controllers.ShowThreads)
	m.Post("/thread/create", controllers.CreateThread)

	m.Run()
}

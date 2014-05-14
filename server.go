package main

import (
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/m0t0k1ch1/martini-sample/controllers"
	"github.com/m0t0k1ch1/martini-sample/models"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
)

func DB() martini.Handler {
	return func(c martini.Context) {
		db, err := models.InitDB("sqlite3", "./data/martini-sample.db", gorp.SqliteDialect{})
		if err != nil {
			panic(err)
		}
		defer db.Db.Close()
		c.Map(db)
		c.Next()
	}
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(DB())

	m.Get("/", controllers.ShowIndex)
	m.Get("/thread/list", controllers.ShowThreads)
	m.Post("/thread/create", controllers.CreateThread)
	m.Get("/thread/:thread_id", controllers.ShowThread)
	m.Post("/post/create", controllers.CreatePost)

	m.Run()
}

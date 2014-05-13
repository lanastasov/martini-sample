package controllers

import (
	"github.com/m0t0k1ch1/martini-sample/models"
	"github.com/martini-contrib/render"
	"net/http"
)

func ShowThreads(r render.Render, db *models.DB) {
	threads, err := db.GetAllThreads()
	if err != nil {
		panic(err)
	}

	r.HTML(200, "thread/list", struct {
		Threads []*models.Thread
	}{
		Threads: threads,
	})
}

func CreateThread(req *http.Request, r render.Render, db *models.DB) {
	title := req.FormValue("title")
	thread := models.NewThread(title)
	if err := db.CreateThread(thread); err != nil {
		panic(err)
	}

	r.Redirect("/thread/list", 302)
}

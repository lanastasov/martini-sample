package controllers

import (
	"github.com/go-martini/martini"
	"github.com/m0t0k1ch1/martini-sample/models"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
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

func ShowThread(r render.Render, db *models.DB, params martini.Params) {
	threadId, err := strconv.Atoi(params["thread_id"])
	if err != nil {
		panic(err)
	}

	thread, err := db.GetThread(threadId)
	if err != nil {
		panic(err)
	}

	posts, err := db.GetPostsByThreadId(threadId)
	if err != nil {
		panic(err)
	}

	r.HTML(200, "thread/index", struct {
		Thread *models.Thread
		Posts  []*models.Post
	}{
		Thread: thread,
		Posts:  posts,
	})
}

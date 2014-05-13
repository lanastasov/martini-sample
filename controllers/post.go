package controllers

import (
	"github.com/m0t0k1ch1/martini-sample/models"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

func CreatePost(req *http.Request, r render.Render, db *models.DB) {
	threadIdStr := req.FormValue("thread_id")
	threadId, err := strconv.Atoi(threadIdStr)
	if err != nil {
		panic(err)
	}
	body := req.FormValue("body")

	post := models.NewPost(threadId, body)
	if err := db.CreatePost(post); err != nil {
		panic(err)
	}

	r.Redirect("/thread/"+threadIdStr, 302)
}

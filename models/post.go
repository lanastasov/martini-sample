package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Post struct {
	Id        int       `db:"id"`
	ThreadId  int       `db:"thread_id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewPost(threadId int, body string) *Post {
	return &Post{
		ThreadId: threadId,
		Body:     body,
	}
}

func (this *Post) PreInsert(s gorp.SqlExecutor) error {
	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	return nil
}

func (this *Post) PreUpdate(s gorp.SqlExecutor) error {
	this.UpdatedAt = time.Now()
	return nil
}

func (db *DB) CreatePost(post *Post) error {
	return db.Insert(post)
}

func (db *DB) GetPostsByThreadId(threadId int) ([]*Post, error) {
	var posts []*Post
	if _, err := db.Select(&posts, "SELECT * FROM post WHERE thread_id = ?", threadId); err != nil {
		return nil, err
	}
	return posts, nil
}

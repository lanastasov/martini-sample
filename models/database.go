package models

import (
	"github.com/naoina/genmai"
)

type DB struct {
	*genmai.DB
}

func InitDB() (*DB, error) {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "data/martini-sample.db")
	if err != nil {
		return nil, err
	}
	db.CreateTableIfNotExists(&Thread{})
	return &DB{db}, nil
}

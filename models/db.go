package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
)

type DB struct {
	*gorp.DbMap
}

func InitDB(driver string, source string, dialect gorp.Dialect) (*DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{
		Db:      db,
		Dialect: dialect,
	}
	dbMap.AddTableWithName(Thread{}, "thread").SetKeys(true, "Id")
	dbMap.AddTableWithName(Post{}, "post").SetKeys(true, "Id")
	if err := dbMap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}
	return &DB{dbMap}, nil
}

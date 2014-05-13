package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
)

type DB struct {
	*gorp.DbMap
}

func InitDB(driver string, source string, dialect gorp.Dialect) (*DB, error) {
	db, err := sql.Open("sqlite3", "data/martini-sample.db")
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: dialect,
	}
	dbmap.AddTableWithName(Thread{}, "thread").SetKeys(true, "Id")
	if err := dbmap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}
	return &DB{dbmap}, nil
}

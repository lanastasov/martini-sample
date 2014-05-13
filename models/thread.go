package models

type Thread struct {
	Id    int `db:"pk"`
	Title string
}

func NewThread(title string) *Thread {
	return &Thread{
		Title: title,
	}
}

func (db *DB) CreateThread(thread *Thread) error {
	_, err := db.Insert(thread)
	return err
}

func (db *DB) GetAllThreads() ([]Thread, error) {
	var threads []Thread
	if err := db.Select(&threads); err != nil {
		return nil, err
	}
	return threads, nil
}

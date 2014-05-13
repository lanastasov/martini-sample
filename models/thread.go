package models

type Thread struct {
	Id    int
	Title string
}

func NewThread(title string) *Thread {
	return &Thread{
		Title: title,
	}
}

func (db *DB) CreateThread(thread *Thread) error {
	return db.Insert(thread)
}

func (db *DB) GetAllThreads() ([]*Thread, error) {
	var threads []*Thread
	if _, err := db.Select(&threads, "SELECT * FROM thread"); err != nil {
		return nil, err
	}
	return threads, nil
}

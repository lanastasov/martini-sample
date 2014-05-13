package models

type Thread struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
}

func NewThread(title string) *Thread {
	return &Thread{
		Title: title,
	}
}

func (db *DB) CreateThread(thread *Thread) error {
	return db.Insert(thread)
}

func (db *DB) GetThread(threadId int) (*Thread, error) {
	thread, err := db.Get(Thread{}, threadId)
	if err != nil || thread == nil {
		return nil, err
	}
	return thread.(*Thread), nil
}

func (db *DB) GetAllThreads() ([]*Thread, error) {
	var threads []*Thread
	if _, err := db.Select(&threads, "SELECT * FROM thread"); err != nil {
		return nil, err
	}
	return threads, nil
}

package models

type Thread struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
}

func (this *Thread) Posts(db *DB) ([]*Post, error) {
	return GetPostsByThreadId(db, this.Id)
}

func NewThread(title string) *Thread {
	return &Thread{
		Title: title,
	}
}

func CreateThread(db *DB, thread *Thread) error {
	return db.Insert(thread)
}

func GetThread(db *DB, threadId int) (*Thread, error) {
	thread, err := db.Get(Thread{}, threadId)
	if err != nil || thread == nil {
		return nil, err
	}
	return thread.(*Thread), nil
}

func GetAllThreads(db *DB) ([]*Thread, error) {
	var threads []*Thread
	if _, err := db.Select(&threads, "SELECT * FROM thread"); err != nil {
		return nil, err
	}
	return threads, nil
}

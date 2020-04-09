package modles

import "time"

type BookList struct {
	Id int `json:"id"`
	userId int `json:"user_id"`
	Name string `json:"name"`
	Ct time.Time `json:"ct"`
	Ut time.Time `json:"ut"`
	bookIds []int `json:"book_ids"`
}
package modles

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Book struct {
	Author string `json:author`
	Name string `json:name`
	BookCover string `json:book_cover`
	Details string `json:details`
	Ct time.Time `json:ct`
	Ut time.Time `json:ut`
}

func (book *Book) AddBook(db *gorm.DB) {
	nowTime := time.Now().UTC()
	db.Create(&Book{
		Author: book.Author,
		Name: book.Name,
		BookCover: book.BookCover,
		Details: book.Details,
		Ct: nowTime,
		Ut: nowTime,
	})
}

func (book *Book) FindBlurryBooks(db *gorm.DB, keyword string) (result []Book, err []error){
	keywordStr := "%%" + keyword + "%%"
	err = db.Where("author like ?", keywordStr).Where("name like ?", keywordStr).Find(&result).GetErrors()
	return
}


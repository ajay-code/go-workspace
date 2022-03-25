package models

import (
	"time"

	"github.com/ajay-code/go-bookstore/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

type Book struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}



func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func (b *Book) Save() *Book {
	db.Save(b)
	return b
}

func GetAllBooks()  []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB){
	var book Book
	db := db.First(&book, id)
	return &book, db
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Delete(book, id)
	return book
}
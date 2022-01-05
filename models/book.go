package models

import (
	"crud-golang/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	Name   string `gorm:""json:"name"`
	Author string `gorm:""json:"author"`
	Status int    `gorm:""json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("id=?", Id).Delete(book)
	return book
}

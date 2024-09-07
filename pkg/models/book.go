package models

import (
	"gorm.io/gorm"
	"github.com/moheddine-belhaj/Book_store/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id uint) (*Book, *gorm.DB) {
	var book Book
	result := db.First(&book, id)
	return &book, result
}

func DeleteBook(id uint) Book {
	var book Book
	db.Delete(&book, id)
	return book
}

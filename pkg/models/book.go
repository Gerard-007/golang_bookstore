package models

import (
	"fmt"
	"github.com/Gerard-007/golang_bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Author      string `json:"author" gorm:"not null"`
	Publication string `json:"publication" gorm:"not null"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	if err := db.AutoMigrate(&Book{}).Error; err != nil {
		_ = fmt.Errorf("")
	}
	_ = fmt.Errorf("")
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", id).Find(&book)
	if db.Error != nil {
		return nil, db
	}
	return &book, db
}

func DeleteBook(id int64) *Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return &book
}

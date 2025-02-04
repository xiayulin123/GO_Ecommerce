package models

import (
	"github.com/jinzhu/gorm"
	"github.com/xiayulin123/Go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:"name"`
	Author      string `gorm:"json:"author"`
	Publication string `gorm:"json:"publication"`
}

func init() {
	config.Connection()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var Book Book
	db.Where("ID=?", Id).Delete(&Book)
	return Book
}

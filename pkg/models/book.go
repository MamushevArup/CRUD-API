package models

import (
	"bookManageGO/pkg/config"
	"gorm.io/gorm"
)

// here we work with db and make CRUD operation using db

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
	var book []Book
	db.Find(&book)
	return book
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(Id int64) Book {
	var del Book
	db.Where("ID=?", Id).Delete(del)
	return del
}
func Delete() *gorm.DB{
	var book []Book
	db.Find(&book).Delete(&book)
	return db
}
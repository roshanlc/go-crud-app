package models

import (
	"github.com/jinzhu/gorm"
	"github.com/roshanlc/go-crud-app/pkg/config"
)

// global Database variable through which all db things happen
var DB *gorm.DB

// Book model to depict a table in mysql db
// Table name, automatically generated by gorm, most likely as : "books"
type Book struct {
	//gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Setup database connection
// and create or update table columns as necessary
func init() {
	config.Connect()
	DB = config.GETDB()
	DB.AutoMigrate(&Book{})
}

// CreateBook creates a new book resource
func (b *Book) CreateBook() *Book {
	DB.NewRecord(b)
	DB.Create(&b)
	return b
}

// GetAllBooks returns all book resources
func GetAllBooks() []Book {

	var Books []Book
	DB.Find(&Books)
	return Books
}

// GetBookByID returns a specific book's details
func GetBookByID(id int) (*Book, *gorm.DB) {

	var getBook Book
	db := DB.Where("ID=?", id).Find(&getBook)
	return &getBook, db

}

// DeleteBook deletes the specified book resource if it exists
func DeleteBook(id int) Book {

	var deletedBook Book

	DB.Where("ID=?", id).Delete(deletedBook)
	return deletedBook
}

package repository

import (
	"book-api/internal/infrastructure/database"

	"book-api/internal/models"
)

var db = database.Conectar()

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := db.Preload("Author").First(&book, id)

	if result.Error != nil {
		return models.Book{}, result.Error
	}

	return book, nil
}

func CreateBook(book models.Book) (models.Book, error) {
	result := db.Create(&book)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

func UpdateBook(book models.Book) (models.Book, error) {
	result := db.Save(&book)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

func DeleteBook(id uint) error {
	result := db.Delete(&models.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

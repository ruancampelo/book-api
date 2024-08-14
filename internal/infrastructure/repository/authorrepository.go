package repository

import (
	"book-api/internal/models"
)

func GetAuthorByID(id uint) (models.Author, error) {
	var author models.Author
	result := db.First(&author, id)
	return author, result.Error
}

func GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	result := db.Find(&authors)
	return authors, result.Error
}

func CreateAuthor(author models.Author) (models.Author, error) {
	result := db.Create(&author)
	return author, result.Error
}

func UpdateAuthor(author models.Author) (models.Author, error) {
	result := db.Save(&author)
	return author, result.Error
}

func DeleteAuthor(id uint) error {
	result := db.Delete(&models.Author{}, id)
	return result.Error
}

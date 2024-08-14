package handlers

import (
	"book-api/internal/infrastructure/repository"
	"log"
	"net/http"
	"strconv"

	"book-api/internal/models"

	"github.com/gin-gonic/gin"
)

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Retrieve details of a specific book by its ID
// @Tags books
// @Param id path int true "Book ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Book details"
// @Failure 400 {object}  map[string]string "Invalid ID format"
// @Failure 404 {object}  map[string]string "Book not found"
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro("Invalid ID format"))
		return
	}

	book, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Book not found"))
		return
	}

	c.JSON(http.StatusOK, models.CreateResultOk(book))

}

// GetAllBooks godoc
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Produce json
// @Success 200 {array}  map[string]interface{}"List of books"
// @Failure 404 {object} map[string]string "Books not found"
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	book, err := repository.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Book not found"))
		return
	}

	response := models.CreateResultOk(book)
	c.JSON(http.StatusOK, response)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the system
// @Tags books
// @Accept json
// @Produce json
// @Param book body map[string]interface{} true "Book input"
// @Success 201 {object} map[string]interface{} "Created book details"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 422 {object} map[string]string "Unprocessable Entity"
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		AuthorID string `json:"author_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro(err.Error()))
		return
	}

	authorID, err := strconv.ParseUint(input.AuthorID, 10, 32)
	if err != nil {
		log.Fatalf("Invalid AuthorID: %v", err)
	}

	book := models.Book{
		Name:     input.Name,
		AuthorID: uint(authorID),
	}

	response, err := repository.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.CreateResultErro(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, models.CreateResultOk(response))
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update the details of an existing book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body map[string]interface{} true "Book input"
// @Success 200 {object} map[string]interface{} "Updated book details"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro(err.Error()))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro("Invalid ID format"))
		return
	}

	book, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Book not found"))
		return
	}

	book.Name = input.Name

	response, err := repository.UpdateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.CreateResultOk(response))
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Remove a specific book from the system by its ID
// @Tags books
// @Param id path int true "Book ID"
// @Produce json
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro("Invalid ID format"))
		return
	}

	book, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Book not found"))
		return
	}

	if err := repository.DeleteBook(book.Id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.CreateResultErro(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

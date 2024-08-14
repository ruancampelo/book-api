package handlers

import (
	"book-api/internal/infrastructure/repository"
	"net/http"
	"strconv"

	"book-api/internal/models"

	"github.com/gin-gonic/gin"
)

// GetAuthorByID godoc
// @Summary Get an author by ID
// @Description Retrieve details of a specific author by their ID
// @Tags authors
// @Param id path int true "Author ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Author details"
// @Failure 400 {object} string "Invalid ID format"
// @Failure 404 {object} string "Author not found"
// @Router /authors/{id} [get]
func GetAuthorByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro("Invalid ID format"))
		return
	}

	author, err := repository.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Author not found"))
		return
	}

	c.JSON(http.StatusOK, author)
}

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Retrieve a list of all authors
// @Tags authors
// @Produce json
// @Success 200 {array} map[string]interface{} "List of authors"
// @Failure 404 {object} string "Authors not found"
// @Router /authors [get]
func GetAllAuthors(c *gin.Context) {
	authors, err := repository.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Book not found"))
		return
	}

	c.JSON(http.StatusOK, authors)
}

// CreateAuthor godoc
// @Summary Create a new author
// @Description Add a new author to the system
// @Tags authors
// @Accept json
// @Produce json
// @Param author body map[string]interface{} true "Author input"
// @Success 201 {object} map[string]string "Created author details"
// @Failure 400 {object} string "Invalid input"
// @Failure 422 {object} string "Unprocessable Entity"
// @Router /authors [post]
func CreateAuthor(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro(err.Error()))
		return
	}

	author := models.Author{
		Name: input.Name,
	}
	author, err := repository.CreateAuthor(author)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.CreateResultErro(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, models.CreateResultOk(author))
}

// UpdateAuthor godoc
// @Summary Update an existing author
// @Description Update the details of an existing author by their ID
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body map[string]string true "Author input"
// @Success 200 {object} map[string]string "Updated author details"
// @Failure 400 {object} string "Invalid input"
// @Failure 404 {object} string "Author not found"
// @Router /authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
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

	author, err := repository.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Author not found"))
		return
	}

	author.Name = input.Name

	updatedAuthor, err := repository.UpdateAuthor(author)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.CreateResultOk(updatedAuthor))
}

// DeleteAuthor godoc
// @Summary Delete an author by ID
// @Description Remove a specific author from the system by their ID
// @Tags authors
// @Param id path int true "Author ID"
// @Produce json
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Author not found"
// @Router /authors/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResultErro("Invalid ID format"))
		return
	}

	author, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.CreateResultErro("Book not found"))
		return
	}

	if err := repository.DeleteAuthor(author.Id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.CreateResultErro(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

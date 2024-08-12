package controller

import (
	"Library-backend/models"
	"Library-backend/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	if err := bc.bookService.CreateBook(context.Background(), &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})

}

func (bc *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bc.bookService.GetBookByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) GetAllBooks(c *gin.Context) {
	books, err := bc.bookService.GetAllBooks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := bc.bookService.DeleteBook(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.bookService.UpdateBook(context.Background(), &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

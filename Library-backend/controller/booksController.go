package controller

import (
	"Library-backend/models"
	"Library-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	bookService services.BookService
}

func NewUserController(userService services.BookService) *UserController {
	return &UserController{bookService: userService}

}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.Book
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	if err := uc.bookService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

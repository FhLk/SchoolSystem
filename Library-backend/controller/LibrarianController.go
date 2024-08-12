package controller

import (
	"Library-backend/models"
	"Library-backend/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LibrarianController struct {
	librarianService services.LibrarianService
}

func NewLibrarianController(librarianService services.LibrarianService) *LibrarianController {
	return &LibrarianController{librarianService: librarianService}
}

func (lc *LibrarianController) CreateLibrarian(c *gin.Context) {
	var librarian models.Librarian
	if err := c.ShouldBindJSON(&librarian); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	if err := lc.librarianService.CreateLibrarian(context.Background(), &librarian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Librarian created successfully"})
}

func (lc *LibrarianController) GetAllLibrarians(c *gin.Context) {
	librarians, err := lc.librarianService.GetAllLibrarians(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, librarians)
}

func (lc *LibrarianController) GetLibrarianByID(c *gin.Context) {
	id := c.Param("id")
	librarian, err := lc.librarianService.GetLibrarianByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, librarian)
}

func (lc *LibrarianController) UpdateLibrarian(c *gin.Context) {
	var librarian models.Librarian
	if err := c.ShouldBindJSON(&librarian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := lc.librarianService.UpdateLibrarian(context.Background(), &librarian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Librarian updated successfully"})
}

func (lc *LibrarianController) DeleteLibrarian(c *gin.Context) {
	id := c.Param("id")
	err := lc.librarianService.DeleteLibrarian(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Librarian deleted successfully"})
}

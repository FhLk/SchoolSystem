package main

import (
	"Library-backend/controller"
	"Library-backend/db"
	"Library-backend/repositories"
	"Library-backend/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// เชื่อมต่อฐานข้อมูล )
	DB, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	bookRepository := repositories.NewBookRepository(DB)
	bookService := services.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	router := gin.Default()
	router.GET("/books/:id", bookController.GetBookByID)
	router.GET("/books/", bookController.GetAllBooks)
	router.POST("/books/", bookController.CreateBook)
	router.PATCH("/books/:id", bookController.UpdateBook)
	router.PUT("/books/:id", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)

	router.Run("localhost:8080")
}

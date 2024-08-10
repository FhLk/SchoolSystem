package main

import (
	"Library-backend/controller"
	"Library-backend/db"
	"Library-backend/repositories"
	"Library-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {

	bookController := initSetupBook()

	router := gin.Default()
	router.GET("/books/:id", bookController.GetBookByID)
	router.GET("/books/", bookController.GetAllBooks)
	router.POST("/books/", bookController.CreateBook)
	router.PATCH("/books/:id", bookController.UpdateBook)
	router.PUT("/books/:id", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)

	router.Run("localhost:8080")
}

func Conn() *gorm.DB {
	DB, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

func initSetupBook() *controller.BookController {
	// เชื่อมต่อฐานข้อมูล )
	DB := Conn()
	bookRepository := repositories.NewBookRepository(DB)
	bookService := services.NewBookService(bookRepository)
	return controller.NewBookController(bookService)
}

func initSetupLibrarian() {
	//DB := Conn()
	//bookRepository := repositories.NewBookRepository(DB)
	//bookService := services.NewBookService(bookRepository)
	//return controller.NewBookController(bookService)
}

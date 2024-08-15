package main

import (
	"Library-backend/controller"
	"Library-backend/db"
	"Library-backend/middleware"
	"Library-backend/repositories"
	"Library-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {
	bookController := initSetupBook()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/api/book", bookController.GetAllBooks)
	bookRoute := router.Group("/api/book")
	{
		bookController := initSetupBook()
		bookRoute.GET("/:id", bookController.GetBookByID)
		bookRoute.GET("/", bookController.GetAllBooks)
		bookRoute.POST("/", bookController.CreateBook)
		bookRoute.PATCH("/:id", bookController.UpdateBook)
		bookRoute.PUT("/:id", bookController.UpdateBook)
		bookRoute.DELETE("/:id", bookController.DeleteBook)
	}

	librarianRoute := router.Group("/api/librarian")
	{
		librarianController := initSetupLibrarian()
		librarianRoute.POST("/", librarianController.CreateLibrarian)
		librarianRoute.GET("/", librarianController.GetAllLibrarians)
		librarianRoute.GET("/:id", librarianController.GetLibrarianByID)
		librarianRoute.PUT("/:id", librarianController.UpdateLibrarian)
		librarianRoute.PATCH("/:id", librarianController.UpdateLibrarian)
		librarianRoute.DELETE("/:id", librarianController.DeleteLibrarian)
	}

	accountRoute := router.Group("/api/account")
	{
		accountController := initSetupAccount()
		accountRoute.POST("/", accountController.CreateAccount)
		accountRoute.GET("/", accountController.GetAllAccount)
		accountRoute.GET("/:id", accountController.GetAccountByID)
		accountRoute.PUT("/:id", accountController.UpdateAccount)
		accountRoute.PATCH("/:id", accountController.UpdateAccount)
		accountRoute.DELETE("/:id", accountController.DeleteAccount)
	}

	memberRoute := router.Group("/api/member")
	{
		memberController := initSetupMember()
		memberRoute.POST("/", memberController.CreateMember)
		memberRoute.GET("/", memberController.GetAllMember)
		memberRoute.GET("/:id", memberController.GetMemberByID)
		memberRoute.PUT("/:id", memberController.UpdateMember)
		memberRoute.PATCH("/:id", memberController.UpdateMember)
		memberRoute.DELETE("/:id", memberController.DeleteMember)
	}
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

func initSetupLibrarian() *controller.LibrarianController {
	DB := Conn()
	librarianRepository := repositories.NewLibrarianRepository(DB)
	librarianService := services.NewLibrarianService(librarianRepository)
	return controller.NewLibrarianController(librarianService)
}

func initSetupAccount() *controller.AccountController {
	DB := Conn()
	accountRepository := repositories.NewAccountRepository(DB)
	accountService := services.NewAccountService(accountRepository)
	return controller.NewAccountController(accountService)
}

func initSetupMember() *controller.MemberController {
	DB := Conn()
	memberRepository := repositories.NewMemberRepository(DB)
	memberService := services.NewMemberService(memberRepository)
	return controller.NewMemberController(memberService)
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm-tutorial/Controller"
	"gorm-tutorial/Google1"
	"gorm-tutorial/Service"
	"gorm-tutorial/Token"
	"gorm-tutorial/initD"
	"os"
)

func main() {

	err := godotenv.Load()

	serverAddress := os.Getenv("serverAddress")

	db, err := initD.ConnectDB()
	if err != nil {
		log.Info("Failed to connect to the database")
	}

	if db == nil {
		log.Info("DB is nil")
	}

	userService := Service.NewUserService(db)
	bookService := Service.NewBookService(db)
	if userService == nil || bookService == nil {
		log.Info("UserService is nil")
	}

	router := gin.Default()

	userController := Controller.NewUserController(userService)
	bookController := Controller.NewBookController(bookService)
	// User
	router.POST("/users", userController.AddUserHandler)
	router.GET("/getAll", userController.GetAllUser)
	router.POST("/keyword", userController.SearchUserHandler)
	router.POST("/login", userController.Login)

	router.POST("/AddBook", bookController.AddBook)
	router.POST("/SreachBook", bookController.SreachBook)
	router.PUT("/UpdateBook", bookController.UpdateBook)
	router.GET("/google", Google1.GoogleLogin)
	router.GET("/login/oauth2/code/google", Google1.GoogleToken)
	router.Use(Token.TokenAuthMiddleware())
	router.GET("/findByUser/:id", userController.GetUserByIDHandler)
	router.Run(serverAddress)

}

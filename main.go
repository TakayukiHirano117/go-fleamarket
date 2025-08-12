package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()

	itemRouter := r.Group("/items")
	itemRouter.GET("/", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("/", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run("localhost:8080")
}

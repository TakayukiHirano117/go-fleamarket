package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{
			ID:          1,
			Name:        "Item 1",
			Price:       1000,
			Description: "Description 1",
			SoldOut:     false,
		},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r:= gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)

	r.Run("localhost:8080")
}

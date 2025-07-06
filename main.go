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
		{
			ID:          2,
			Name:        "Item 2",
			Price:       2000,
			Description: "Description 2",
			SoldOut:     false,
		},
		{
			ID:          3,
			Name:        "Item 3",
			Price:       3000,
			Description: "Description 3",
			SoldOut:     false,
		},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	router := gin.Default()
	router.GET("/items", itemController.FindAll)
	router.GET("/items/:id", itemController.FindById)
	router.Run("localhost:8080")
}

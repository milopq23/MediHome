package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func InventoryRoute(r *gin.RouterGroup) {
	inventoryRepo := repository.NewInventoryRepository()
	inventoryService := service.NewInventoryService(inventoryRepo)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)
	inventory := r.Group("/admin/inventory")
	{
		inventory.GET("/", inventoryHandler.GetAll)
		inventory.POST("/", inventoryHandler.Create)
		// inventory.PATCH("/:id", inventoryHandler.Patch)
		inventory.DELETE("/:id", inventoryHandler.Delete)
	}
}

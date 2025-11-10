package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func CartRoute(r *gin.Engine) {
	cartRepo := repository.NewCartRepository()
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)
	cart := r.Group("/cart")
	{
		cart.GET("/:id", cartHandler.GetCartUser)
		cart.POST("/:id", cartHandler.AddCart)
		cart.PATCH("/:id", cartHandler.UpdateCart)
		cart.DELETE("/:id", cartHandler.DeleteItemCart)
	}
}

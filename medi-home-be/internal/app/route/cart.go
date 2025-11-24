package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"
	"medi-home-be/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoute(r *gin.RouterGroup) {
	cartRepo := repository.NewCartRepository()
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	userAuth := middleware.AuthorizeMiddleware(middleware.User)

	cart := r.Group("/cart")
	{
		cart.GET("/", userAuth, cartHandler.GetCartUser)
		cart.POST("/", userAuth, cartHandler.AddCart)
		cart.PATCH("/:id", cartHandler.UpdateCart)
		cart.DELETE("/:id", cartHandler.DeleteItemCart)
	}
}

package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func ShippingRoute(r *gin.Engine) {
	shippingRepo := repository.NewShippingRepository()
	shippingService := service.NewShippingService(shippingRepo)
	shippingHandler := handler.NewShippingHandler(shippingService)
	shipping := r.Group("/shipping")
	{
		shipping.GET("/", shippingHandler.GetAll)
		// shipping.GET("/:id", shippingHandler.GetByID)
		// shipping.POST("/", shippingHandler.Create)
		// shipping.PUT("/:id", shippingHandler.Patch)
		// shipping.PATCH("/:id", shippingHandler.Patch)
		shipping.DELETE("/:id", shippingHandler.Delete)
	}
}
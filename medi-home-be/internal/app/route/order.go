package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.RouterGroup) {
	cartRepo := repository.NewCartRepository()
	shippingRepo := repository.NewShippingRepository()
	voucherRepo := repository.NewVoucherRepository()

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo, cartRepo, shippingRepo, voucherRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	order := r.Group("/order")
	{
		order.GET("/", orderHandler.GetAllOrder)
		order.POST("/checkout/:id", orderHandler.CheckOut)
	}
}

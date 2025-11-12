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
	inventoryRepo := repository.NewInventoryRepository()
	logTransaction := repository.NewStockTransactionRepository()

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo, cartRepo, shippingRepo, voucherRepo, inventoryRepo, logTransaction)
	orderHandler := handler.NewOrderHandler(orderService)

	order := r.Group("/order")
	{
		//lấy tất cả và lọc theo loại
		order.GET("/", orderHandler.GetOrders)
		order.GET("/:id", orderHandler.GetDetailOrder)
		order.GET("/user/:id", orderHandler.GetUserOrders)
		order.POST("/checkout/:id", orderHandler.CheckOut)
	}
	// userOrder := r.Group("/order")
	// {
	// 	//lấy tất cả và lọc theo loại
	// 	order.GET("/", orderHandler.GetOrders)
	// 	// order.GET("/:id", orderHandler.GetDetailOrder)
	// 	order.GET("/:id",orderHandler.GetUserOrders)
	// 	order.POST("/checkout/:id", orderHandler.CheckOut)
	// }
}

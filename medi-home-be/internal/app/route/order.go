package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"
	"medi-home-be/pkg/middleware"

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

	userAuth := middleware.AuthorizeMiddleware(middleware.User)

	order := r.Group("/order")
	{
		//lấy tất cả và lọc theo loại
		order.GET("/", orderHandler.GetOrders)
		order.GET("/:id", orderHandler.GetDetailOrder)
		order.GET("/user/", userAuth, orderHandler.GetUserOrders)
		order.PATCH("/approve/:id", orderHandler.ApproveStatus)
		order.POST("/checkout/", userAuth, orderHandler.CheckOut)
	}
}

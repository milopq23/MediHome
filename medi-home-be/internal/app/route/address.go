package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"
	"medi-home-be/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AddrressRoute(r *gin.RouterGroup) {
	addressRepo := repository.NewAddressRepository()
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)

	userAuth := middleware.AuthorizeMiddleware(middleware.User)

	address := r.Group("/address")
	{
		address.GET("/", userAuth, addressHandler.GetAddress)
		address.POST("/", userAuth, addressHandler.AddAddress)
	}
}

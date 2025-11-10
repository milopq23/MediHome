package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func AddrressRoute(r *gin.RouterGroup) {
	addressRepo := repository.NewAddressRepository()
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)
	address := r.Group("/address")
	{
		address.POST("/:id", addressHandler.AddAddress)
	}
}

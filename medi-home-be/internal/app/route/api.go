package route

import (

	"github.com/gin-gonic/gin"
)

func APIRoute() *gin.Engine{
	r := gin.Default()

	UserRoutes(r)
	AdminRoutes(r)
	MedicineRoute(r)
	ShippingRoute(r)
	
	return r
}
package route

import (
	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func APIRoute() *gin.Engine {
	r := gin.Default()

	UserRoutes(r)
	AdminRoutes(r)
	MedicineRoute(r)
	ShippingRoute(r)
	DosageFormRoute(r)

	return r
}

package route

import (
	"github.com/gin-gonic/gin"
)

func APIRoute(r *gin.Engine) {
	api := r.Group("/api")
	UserRoutes(api)
	MedicineRoute(api)

	ShippingRoute(api)
	DosageFormRoute(api)
	UploadRoute(api)
	AdminRoutes(api)
}

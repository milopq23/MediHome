package route

import (
	"github.com/gin-gonic/gin"
)

func APIRoute(r *gin.Engine) {
	api := r.Group("/api")
	UserRoutes(api)
	MedicineRoute(api)
	InventoryRoute(api)
	AddrressRoute(api)
	OrderRoute(api)
	CartRoute(api)
	ShippingRoute(api)
	DosageFormRoute(api)
	VoucherRoutes(api)
	UploadRoute(api)
	AdminRoutes(api)
}

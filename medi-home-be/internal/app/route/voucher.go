package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func VoucherRoutes(r *gin.RouterGroup) {
	voucherRepo := repository.NewVoucherRepository()
	voucherService := service.NewVoucherService(voucherRepo)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	voucher := r.Group("/voucher")
	{
		voucher.GET("/", voucherHandler.GetAllVoucher)
		voucher.GET("/active", voucherHandler.FindActiveVoucher)
		voucher.POST("/", voucherHandler.CreateVoucher)
		voucher.GET("/:code", voucherHandler.GetVoucherByCode)
		voucher.PATCH("/", voucherHandler.PatchVoucher)
		voucher.DELETE("/:id", voucherHandler.DeleteVoucher)
	}
}

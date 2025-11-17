package route

import (
	"log"
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/service"
	"medi-home-be/pkg/util"

	"github.com/gin-gonic/gin"
)

func UploadRoute(r *gin.RouterGroup) {
	cld, err := util.UltiCloudinary()
	if err != nil {
		log.Fatal("Lỗi khi khởi tạo Cloudinary:", err)
	}
	service := service.NewCloudinaryService(cld)
	handler := handler.NewUploadHandler(service)

	r.POST("/single", handler.SingleUpload)
	r.POST("/multi",handler.MultiUpload)
}

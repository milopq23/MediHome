package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func MedicineRoute(r *gin.RouterGroup) {
	medicineRepo := repository.NewMedicineRepository()
	cartRepo := repository.NewCartRepository()
	medicineService := service.NewMedicineService(medicineRepo, cartRepo)
	medicineHandler := handler.NewMedicineHandler(medicineService)
	adminMedicine := r.Group("/admin/medicine")
	{
		adminMedicine.GET("/", medicineHandler.GetAll)
		adminMedicine.GET("/:id", medicineHandler.GetByID)
		adminMedicine.POST("/", medicineHandler.Create)
		adminMedicine.POST("/:medicine_id/images", medicineHandler.SaveImages)
		adminMedicine.PUT("/:id", medicineHandler.Patch)
		adminMedicine.PATCH("/:id", medicineHandler.Patch)
		adminMedicine.DELETE("/:id", medicineHandler.Delete)
	}

	medicine := r.Group("/medicine")
	{
		medicine.GET("/", medicineHandler.ListMedicine)
		// medicine.GET("/:id", medicineHandler.DetailMedicine)
		medicine.GET("/:id", medicineHandler.DetailMedicineUser)
	}

	medicineCateRepo := repository.NewMedicineCateRepository()
	medicineCateService := service.NewMedicineCateService(medicineCateRepo)
	medicineCateHandler := handler.NewMedicineCateHandler(medicineCateService)
	medicineCate := r.Group("/admin/medicinecate")
	{
		medicineCate.GET("/", medicineCateHandler.GetAll)
		medicineCate.GET("/children", medicineCateHandler.ListChildren)
		medicineCate.GET("/parent", medicineCateHandler.ListParent)
		medicineCate.POST("/parent", medicineCateHandler.CreateParentCate)
		medicineCate.POST("/children", medicineCateHandler.Create)
		medicineCate.PUT("/:id", medicineCateHandler.Patch)
		medicineCate.PATCH("/:id", medicineCateHandler.Patch)
		medicineCate.DELETE("/:id", medicineCateHandler.Delete)
	}
}

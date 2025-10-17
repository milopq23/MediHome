package route

import (
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"
	"medi-home-be/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func MedicineRoute(r *gin.RouterGroup) {
	medicineRepo := repository.NewMedicineRepository()
	medicineService := service.NewMedicineService(medicineRepo)
	medicineHandler := handler.NewMedicineHandler(medicineService)
	adminMedicine := r.Group("/admin/medicine")
	{
		adminMedicine.GET("/", medicineHandler.GetAll)
		adminMedicine.GET("/:id", medicineHandler.GetByID)
		adminMedicine.POST("/", medicineHandler.Create)
		adminMedicine.PUT("/:id", medicineHandler.Patch)
		adminMedicine.PATCH("/:id", medicineHandler.Patch)
		adminMedicine.DELETE("/:id", medicineHandler.Delete)
	}
	
	medicine := r.Group("/medicine")
	{
		medicine.GET("/",medicineHandler.ListMedicine)
		medicine.GET("/:id", medicineHandler.DetailMedicine)
	}

	medicineCateRepo := repository.NewMedicineCateRepository()
	medicineCateService := service.NewMedicineCateService(medicineCateRepo)
	medicineCateHandler := handler.NewMedicineCateHandler(medicineCateService)
	medicineCate := r.Group("/medicine-cate")
	{
		medicineCate.GET("/", medicineCateHandler.GetAll)
		medicineCate.GET("/children/:id", medicineCateHandler.ListChildren)
		medicineCate.POST("/parent", medicineCateHandler.CreateParentCate)
		medicineCate.POST("/children", medicineCateHandler.Create)
		medicineCate.PUT("/:id", medicineCateHandler.Patch)
		medicineCate.PATCH("/:id", medicineCateHandler.Patch)
		medicineCate.DELETE("/:id", medicineCateHandler.Delete)
	}
}
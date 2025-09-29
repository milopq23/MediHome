package route

import (
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"
	"medi-home-be/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func MedicineRoute(r *gin.Engine) {
	medicineRepo := repository.NewMedicineRepository()
	medicineService := service.NewMedicineService(medicineRepo)
	medicineHandler := handler.NewMedicineHandler(medicineService)
	medicine := r.Group("/medicine")
	{
		medicine.GET("/", medicineHandler.GetAll)
		medicine.GET("/:id", medicineHandler.GetByID)
		medicine.POST("/", medicineHandler.Create)
		medicine.PUT("/:id", medicineHandler.Patch)
		medicine.PATCH("/:id", medicineHandler.Patch)
		medicine.DELETE("/:id", medicineHandler.Delete)


	}

	medicineCateRepo := repository.NewMedicineCateRepository()
	medicineCateService := service.NewMedicineCateService(medicineCateRepo)
	medicineCateHandler := handler.NewMedicineCateHandler(medicineCateService)
	medicineCate := r.Group("/medicine-cate")
	{
		medicineCate.GET("/", medicineCateHandler.GetAll)
		medicineCate.GET("/:id/children", medicineCateHandler.ListChildren)
		medicineCate.POST("/", medicineCateHandler.Create)
		medicineCate.PUT("/:id", medicineCateHandler.Patch)
		medicineCate.PATCH("/:id", medicineCateHandler.Patch)
		medicineCate.DELETE("/:id", medicineCateHandler.Delete)
	}
}
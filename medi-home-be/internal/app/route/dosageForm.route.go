package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func DosageFormRoute(r *gin.Engine){
	dosageRepo := repository.NewDosageFormRepository()
	dosageService := service.NewDosageFormService(dosageRepo)
	dosageHandler := handler.NewDosageFormHandler(dosageService)

	dosage := r.Group("/dosage")
	{
		dosage.GET("/:id",dosageHandler.GetByID)
		dosage.POST("/",dosageHandler.Create)
		dosage.PATCH("/:id",dosageHandler.Patch)
		dosage.DELETE("/:id",dosageHandler.Delete)

	}
}
package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup){
	adminRepo := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)
	user := r.Group("/admin")
	{
		user.GET("/",adminHandler.GetAll)
		user.GET("/:id",adminHandler.GetByID)
		user.GET("/total",adminHandler.TotalAdmin)
		// user.GET("/total_admin",adminHandler.TotalAdmin)
		user.POST("/",adminHandler.Create)
		// user.PUT("/update",userHandler.Update)
		user.PATCH("/:id",adminHandler.Patch)
		user.DELETE("/:id",adminHandler.Delete)
	}
}
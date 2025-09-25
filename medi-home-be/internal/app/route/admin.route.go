package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine){
	adminRepo := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminService)
	user := r.Group("/user")
	{
		user.GET("/",adminHandler.GetAll)
		user.GET("/:id",adminHandler.GetByID)
		user.POST("/create",adminHandler.Create)
		// user.PUT("/update",userHandler.Update)
		user.PATCH("/update/:id",adminHandler.Patch)
		user.DELETE("/delete",adminHandler.Delete)
	}
}
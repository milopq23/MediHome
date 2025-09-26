package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){

	// 
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	login := r.Group("/")
	{
		login.POST("/login",userHandler.Login)
		login.POST("/register",userHandler.Register)
	}
	user := r.Group("/user")
	{
		user.GET("/",userHandler.GetAll)
		user.GET("/:id",userHandler.GetByID)
		user.POST("/create",userHandler.Create)
		// user.PUT("/update",userHandler.Update)
		user.PATCH("/update/:id",userHandler.Patch)
		user.DELETE("/delete",userHandler.Delete)
	}
}
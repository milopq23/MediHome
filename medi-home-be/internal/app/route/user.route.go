package route

import (
	"log"
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	log.Print("user")
	// login := r.Group("/")
	// {
	// 	login.POST("/login",userHandler.Login)
	// 	login.POST("/register",userHandler.Register)
	// }
	user := r.Group("/api/user")
	{
		log.Print("api user")
		user.GET("/",userHandler.GetAll)
		// user.GET("/:id",userHandler.GetByID)
		// user.POST("/",userHandler.Create)
		// // user.PUT("/update",userHandler.Update)
		// user.PATCH("/:id",userHandler.Patch)
		// user.DELETE("/:id",userHandler.Delete)
	}
}
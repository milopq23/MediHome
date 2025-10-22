package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	login := r.Group("/")
	// login.Use(middleware.Authorize(middleware.User))
	{
		login.POST("/login", userHandler.Login)
		login.POST("/register", userHandler.Register)
	}
	user := r.Group("/user")
	// user.Use(middleware.AdminAuthorize(middleware.Admin))
	{
		user.GET("/", userHandler.GetAll)
		user.GET("/total", userHandler.TotalActive)
		user.GET("/:id", userHandler.GetByID)
		user.POST("/", userHandler.Create)
		// // user.PUT("/update",userHandler.Update)
		user.PATCH("/:id", userHandler.Patch)
		user.DELETE("/:id", userHandler.Delete)
	}
}

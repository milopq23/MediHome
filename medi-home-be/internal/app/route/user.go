package route

import (
	"medi-home-be/internal/app/handler"
	"medi-home-be/internal/app/repository"
	"medi-home-be/internal/app/service"
	"medi-home-be/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userAuth := middleware.AuthorizeMiddleware(middleware.User)

	user := r.Group("/")
	{
		user.POST("/login", userAuth, userHandler.Login)
		user.POST("/register", userHandler.Register)
		user.GET("/profile/:id", userHandler.Profile)
		user.PATCH("/profile/", userAuth, userHandler.Patch)
		user.POST("/logout", userHandler.LogOut)
	}
	userAdmin := r.Group("/user")
	{

		userAdmin.GET("/", userHandler.GetAll)
		userAdmin.GET("/total", userHandler.TotalActive)
		userAdmin.GET("/:id", userHandler.GetByID)
		userAdmin.POST("/", userHandler.Create)
		userAdmin.PATCH("/:id", userHandler.Patch)
		userAdmin.DELETE("/:id", userHandler.Delete)
	}
}

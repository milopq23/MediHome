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

	// userAuth := middleware.AuthorizeMiddleware(middleware.User)

	user := r.Group("/")
	{
		user.POST("/login", userHandler.Login)
		user.POST("/register", userHandler.Register)
		user.GET("/profile/", userHandler.Profile)
		user.PATCH("/profile/", userHandler.Patch)
		user.POST("/logout", userHandler.LogOut)
	}
	userAdmin := r.Group("/admin/user")
	{

		userAdmin.GET("/", userHandler.GetAll)
		userAdmin.GET("/total", userHandler.TotalActive)
		userAdmin.GET("/:id", userHandler.GetByID)
		userAdmin.POST("/", userHandler.Create)
		userAdmin.PATCH("/:id", userHandler.Patch)
		userAdmin.DELETE("/:id", userHandler.Delete)
	}

	cartRepo := repository.NewCartRepository()
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)
	cart := r.Group("/cart")
	{
		cart.GET("/:id", cartHandler.GetCartUser)
		cart.POST("/:id", cartHandler.AddCart)
		cart.PATCH("/:id",cartHandler.UpdateCart)
	}

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo,cartRepo)
	orderHandler := handler.NewOrderHandler(orderService)
	order := r.Group("/order")
	{
		order.POST("/checkout/:id",orderHandler.CheckOut)
	}

	addressRepo := repository.NewAddressRepository()
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)
	address := r.Group("/address")
	{
		address.POST("/:id", addressHandler.AddAddress)
	}
}

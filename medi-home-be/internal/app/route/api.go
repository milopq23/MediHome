package route

import (

	"github.com/gin-gonic/gin"
)

func APIRoute() *gin.Engine{
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/",GetUser)
		userGroup.POST("/create",CreateUser)
		userGroup.GET("/id",GetUserByID)
		userGroup.PUT("/update",UpdateUser)
		userGroup.DELETE("/delete",DeleteUser)
	}

}
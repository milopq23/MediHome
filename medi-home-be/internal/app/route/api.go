package route

import (

	"github.com/gin-gonic/gin"
)

func APIRoute() *gin.Engine{
	r := gin.Default()

	UserRoutes(r)
	AdminRoutes(r)
	MedicineRoute(r)
	
	return r
}
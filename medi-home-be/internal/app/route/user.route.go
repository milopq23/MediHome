package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c*gin.Engine){
	user := userService.GetAllUser()
	c.JSON(http.StatusOK,user)
}
package main

import (
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/route"
	"net/http"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	config.ConnectDB()
	r := route.APIRoute()	

	if err := r.Run(":8000"); err != nil{
		log.Fatal("Failed",err)
	}
}

func Hello(c*gin.Context) {
	hello := gin.H{
		"message":"HelloWorld",
	}
	c.JSON(http.StatusOK,hello)
}
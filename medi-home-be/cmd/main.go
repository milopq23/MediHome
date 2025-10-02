package main

import (
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/route"
	"net/http"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.ConnectDB()

	r := route.APIRoute()
	//cors.Default()
	r.Use(CORSMiddleware())

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed", err)
	}
}

func Hello(c *gin.Context) {
	hello := gin.H{
		"message": "HelloWorld",
	}
	c.JSON(http.StatusOK, hello)
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

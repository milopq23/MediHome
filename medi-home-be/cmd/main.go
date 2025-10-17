package main

import (
	"log"
	"medi-home-be/config"
	"medi-home-be/internal/app/route"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	route.APIRoute(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed", err)
	}

}

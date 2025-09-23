package main

import (
	"log"
	"medi-home-be/config"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main(){
	config.ConnectDB()
	r := gin.Default()

	

	

	if err := r.Run(":8000"); err != nil{
		log.Fatal("Failed",err)
	}
}
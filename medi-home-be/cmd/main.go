package main

import(
	"log"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main(){
	r := gin.Default()

	
	if err := r.Run(":8000"); err != nil{
		log.Fatal("Failed",err)
	}
}
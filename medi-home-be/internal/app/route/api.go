package route

import (
	// "fmt"
	// "time"

	// "github.com/gin-contrib/cors"
	"log"

	"github.com/gin-gonic/gin"
)

func APIRoute(r *gin.Engine) {

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// 	 AllowOriginFunc: func(origin string) bool {
	//     log.Println("CORS check origin:", origin)
	//     return origin == "http://localhost:3000"
	// },
	// }))
	log.Print("api duoc")

	UserRoutes(r)
	AdminRoutes(r)
	MedicineRoute(r)
	ShippingRoute(r)
	DosageFormRoute(r)
}

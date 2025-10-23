package middleware

import (
	"net/http"

	"medi-home-be/pkg/util"

	"github.com/gin-gonic/gin"
)

type Permission string

const (
	Admin Permission = "admin"
	Staff Permission = "staff"
	User  Permission = "user"
)

// Authorize middleware kiểm tra token và quyền user
// func AdminAuthorize(required Permission) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
// 			return
// 		}

// 		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

// 		claims, err := util.ParseJWT(tokenStr)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			return
// 		}

// 		if claims.Role != string(required) {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
// 			return
// 		}

// 		// Lưu thông tin claims vào context Gin để handler có thể lấy
// 		c.Set("claims", claims)

// 		c.Next()
// 	}
// }

func UserAuthorize(required Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token")
		if err != nil || tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		claims, err := util.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims.Role != string(required) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		// Lưu thông tin claims vào context Gin để handler có thể lấy
		c.Set("claims", claims)

		c.Next()
	}
}

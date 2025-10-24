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

var rolePriority = map[Permission]int{
	User:  1,
	Staff: 2,
	Admin: 3,
}

func parseTokenAndSetClaims(c *gin.Context) (*util.Claims, bool) {
	tokenStr, err := c.Cookie("token")
	if err != nil || tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return nil, false
	}

	claims, err := util.ParseJWT(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return nil, false
	}

	// Lưu thông tin claims vào context Gin để handler có thể lấy
	c.Set("claims", claims)
	c.Set("role", claims.Role)
	return claims, true
}

func AuthorizeMiddleware(minimumRole Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := parseTokenAndSetClaims(c)
		if !ok {
			return
		}

		userRole := Permission(claims.Role)

		userLevel, ok1 := rolePriority[userRole]
		minLevel, ok2 := rolePriority[minimumRole]

		if !ok1 || !ok2 || userLevel < minLevel {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		c.Next()
	}
}

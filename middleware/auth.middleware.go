package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/helpers"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("authToken")
		result, err := helpers.ValidateJWT(authToken, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized!"})
			c.Abort()
		}

		c.Set("userId", result.UserId)

		c.Next()
	}
}

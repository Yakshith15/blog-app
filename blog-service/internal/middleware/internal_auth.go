package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InternalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		expectedToken := os.Getenv("INTERNAL_SERVICE_TOKEN")
		if expectedToken == "" {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		token := c.GetHeader("X-Internal-Token")
		if token == "" || token != expectedToken {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

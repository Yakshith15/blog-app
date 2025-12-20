package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/Yakshith15/blog-app/blog-service/internal/model"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UNAUTHORIZED",
				"message": "Missing or invalid Authorization header",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Ensure token algorithm is what we expect
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "INVALID_TOKEN",
				"message": "Token is invalid or expired",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "INVALID_TOKEN",
				"message": "Invalid token claims",
			})
			return
		}

		// Extract userId
		sub, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "INVALID_TOKEN",
				"message": "Missing subject in token",
			})
			return
		}

		userID, err := uuid.Parse(sub)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "INVALID_TOKEN",
				"message": "Invalid user ID in token",
			})
			return
		}

		// Extract emailVerified
		emailVerified, ok := claims["emailVerified"].(bool)
		if !ok {
			// default to false if missing (safe)
			emailVerified = false
		}

		// Store auth context
		c.Set("auth", model.AuthContext{
			UserID:        userID,
			EmailVerified: emailVerified,
		})

		c.Next()
	}
}

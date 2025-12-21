package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	port := os.Getenv("PORT") 
	if port == "" {
		port = "8083"
	}

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
	router.Run(":" + port)
	
}
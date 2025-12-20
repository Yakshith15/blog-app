package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	router := gin.Default()
	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	log.Println("Starting Blog Service on port", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
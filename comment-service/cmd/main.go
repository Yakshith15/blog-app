package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)


func main() {
	port := os.Getenv("PORT") 
	if port == "" {
		port = "8083"
	}

	blogServiceURL := os.Getenv("BLOG_SERVICE_URL")
	if blogServiceURL == "" {
		log.Fatal("BLOG_SERVICE_URL is not set")
	}

	internalToken := os.Getenv("INTERNAL_SERVICE_TOKEN")
	if internalToken == "" {
		log.Fatal("INTERNAL_SERVICE_TOKEN is not set")
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// db := config.NewDB()


}
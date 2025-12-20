package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Yakshith15/blog-app/blog-service/internal/handler"
	"github.com/Yakshith15/blog-app/blog-service/internal/middleware"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	router := gin.Default()

	router.Use(middleware.JWTAuthMiddleware())
	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	blogHandler := handler.NewBlogHandler()

	router.GET("/blogs", blogHandler.GetBlogs)
	router.GET("/blogs/:id", blogHandler.GetBlogByID)


	log.Println("Starting Blog Service on port", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
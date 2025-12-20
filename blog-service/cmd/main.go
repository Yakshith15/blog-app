package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Yakshith15/blog-app/blog-service/internal/handler"
	"github.com/Yakshith15/blog-app/blog-service/internal/middleware"
	"github.com/Yakshith15/blog-app/blog-service/internal/config"
	"github.com/Yakshith15/blog-app/blog-service/internal/repository"
	"github.com/Yakshith15/blog-app/blog-service/internal/service"
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

	db := config.NewDB()

	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogHandler := handler.NewBlogHandler(blogService)
	
	router.GET("/blogs", blogHandler.GetBlogs)
	router.GET("/blogs/:id", blogHandler.GetBlogByID)
	router.POST("/blogs", blogHandler.CreateBlog)
	router.PUT("/blogs/:id", blogHandler.UpdateBlog)

	log.Println("Starting Blog Service on port", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
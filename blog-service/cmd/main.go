package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yakshith15/blog-app/blog-service/internal/config"
	"github.com/Yakshith15/blog-app/blog-service/internal/handler"
	"github.com/Yakshith15/blog-app/blog-service/internal/middleware"
	"github.com/Yakshith15/blog-app/blog-service/internal/repository"
	"github.com/Yakshith15/blog-app/blog-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	router := gin.Default()

	// Public (no auth)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	db := config.NewDB()

	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)

	blogHandler := handler.NewBlogHandler(blogService)
	internalBlogHandler := handler.NewInternalBlogHandler(blogService)

	api := router.Group("/")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/blogs", blogHandler.GetBlogs)
		api.GET("/blogs/:id", blogHandler.GetBlogByID)
		api.POST("/blogs", blogHandler.CreateBlog)
		api.PUT("/blogs/:id", blogHandler.UpdateBlog)
		api.DELETE("/blogs/:id", blogHandler.DeleteBlog)
	}

	internal := router.Group("/internal")
	internal.Use(middleware.InternalAuthMiddleware())
	{
		internal.GET("/blogs/:id", internalBlogHandler.CheckBlogExists)
	}

	log.Println("Starting Blog Service on port", port)
	log.Fatal(router.Run(":" + port))
}

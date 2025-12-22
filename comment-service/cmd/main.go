package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Yakshith15/blog-app/comment-service/internal/client"
	"github.com/Yakshith15/blog-app/comment-service/internal/config"
	"github.com/Yakshith15/blog-app/comment-service/internal/handler"
	"github.com/Yakshith15/blog-app/comment-service/internal/middleware"
	"github.com/Yakshith15/blog-app/comment-service/internal/repository"
	"github.com/Yakshith15/blog-app/comment-service/internal/service"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
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

	db := config.NewDB()

	blogClient := client.NewBlogClient(
		blogServiceURL,
		internalToken,
		5*time.Second,
	)

	commentRepo := repository.NewCommentRepository(db)

	commentService := service.NewCommentService(
		commentRepo,
		blogClient,
	)

	commentHandler := handler.NewCommentHandler(commentService)

	api := router.Group("/")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/blogs/:blogId/comments", commentHandler.GetComments)
		api.POST("/blogs/:blogId/comments", commentHandler.CreateComment)
		api.DELETE("/comments/:id", commentHandler.DeleteComment)
	}

	log.Println("Starting Comment Service on port", port)
	log.Fatal(router.Run(":" + port))
}

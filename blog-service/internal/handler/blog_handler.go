package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Yakshith15/blog-app/blog-service/internal/model"
)

type BlogHandler struct {
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

// GET /blogs
func (h *BlogHandler) GetBlogs(c *gin.Context) {

	blogs := []model.Blog{
		{
			ID:        "1",
			AuthorID:  "user-1",
			Title:     "First Blog",
			Content:   "This is the first blog content",
			CreatedAt: time.Now(),
		},
		{
			ID:        "2",
			AuthorID:  "user-2",
			Title:     "Second Blog",
			Content:   "This is the second blog content",
			CreatedAt: time.Now(),
		},
	}

	c.JSON(http.StatusOK, blogs)
}

// GET /blogs/:id
func (h *BlogHandler) GetBlogByID(c *gin.Context) {

	id := c.Param("id")

	blog := model.Blog{
		ID:        id,
		AuthorID:  "user-1",
		Title:     "Sample Blog",
		Content:   "This is a sample blog",
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, blog)
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Yakshith15/blog-app/blog-service/internal/model"
	"github.com/Yakshith15/blog-app/blog-service/internal/service"
)

type BlogHandler struct {
	service *service.BlogService
}

func NewBlogHandler(service *service.BlogService) *BlogHandler {
	return &BlogHandler{service: service}
}

func (h *BlogHandler) GetBlogs(c *gin.Context) {

	blogs, err := h.service.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to fetch blogs",
		})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) GetBlogByID(c *gin.Context) {

	blogID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "INVALID_ID",
			"message": "Invalid blog ID",
		})
		return
	}

	blog, err := h.service.GetByID(blogID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "NOT_FOUND",
			"message": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {

	authCtx, exists := c.Get("auth")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "UNAUTHORIZED",
			"message": "Authentication required",
		})
		return
	}

	auth := authCtx.(model.AuthContext)

	if !auth.EmailVerified {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error":   "EMAIL_NOT_VERIFIED",
			"message": "Email verification required to create blogs",
		})
		return
	}

	var req model.CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": "Invalid request body",
		})
		return
	}

	blog, err := h.service.CreateBlog(
		auth.UserID,
		req.Title,
		req.Content,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to create blog",
		})
		return
	}

	c.JSON(http.StatusCreated, blog)
}


func (h *BlogHandler) UpdateBlog(c *gin.Context) {

	authCtx, exists := c.Get("auth")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "UNAUTHORIZED",
			"message": "Authentication required",
		})
		return
	}

	auth := authCtx.(model.AuthContext)

	if !auth.EmailVerified {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error":   "EMAIL_NOT_VERIFIED",
			"message": "Email verification required to update blogs",
		})
		return
	}

	blogID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "INVALID_ID",
			"message": "Invalid blog ID",
		})
		return
	}

	var req model.UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": "Invalid request body",
		})
		return
	}

	updatedBlog, err := h.service.UpdateBlog(
		blogID,
		auth.UserID,
		req.Title,
		req.Content,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error":   "FORBIDDEN",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

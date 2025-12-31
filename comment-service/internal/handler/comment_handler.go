package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Yakshith15/blog-app/comment-service/internal/model"
	"github.com/Yakshith15/blog-app/comment-service/internal/service"
)

type CommentHandler struct {
	service *service.CommentService
}

func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// GET /blogs/:blogId/comments
func (h *CommentHandler) GetComments(c *gin.Context) {

	blogID, err := uuid.Parse(c.Param("blogId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "INVALID_BLOG_ID",
			"message": "Invalid blog ID",
		})
		return
	}

	comments, err := h.service.GetCommentsByBlogID(blogID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to fetch comments",
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// POST /blogs/:blogId/comments
func (h *CommentHandler) CreateComment(c *gin.Context) {

	// ---- Auth context ----
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
			"message": "Email verification required to comment",
		})
		return
	}

	// ---- Blog ID ----
	blogID, err := uuid.Parse(c.Param("blogId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "INVALID_BLOG_ID",
			"message": "Invalid blog ID",
		})
		return
	}

	// ---- Request body ----
	var req model.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BAD_REQUEST",
			"message": "Invalid request body",
		})
		return
	}

	comment, err := h.service.CreateComment(blogID, auth.UserID, req.Content)
	if err != nil {

		if errors.Is(err, service.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "BLOG_NOT_FOUND",
				"message": "Blog does not exist",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to create comment",
		})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// DELETE /comments/:id
func (h *CommentHandler) DeleteComment(c *gin.Context) {

	// ---- Auth context ----
	authCtx, exists := c.Get("auth")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "UNAUTHORIZED",
			"message": "Authentication required",
		})
		return
	}

	auth := authCtx.(model.AuthContext)

	// ---- Comment ID ----
	commentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "INVALID_COMMENT_ID",
			"message": "Invalid comment ID",
		})
		return
	}

	// ---- Blog ID (required for ownership enforcement) ----
	blogIDParam := c.Query("blogId")
	if blogIDParam == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "BLOG_ID_REQUIRED",
			"message": "blogId query param is required",
		})
		return
	}

	blogID, err := uuid.Parse(blogIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "INVALID_BLOG_ID",
			"message": "Invalid blog ID",
		})
		return
	}

	err = h.service.DeleteComment(commentID, auth.UserID, blogID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error":   "FORBIDDEN",
			"message": "Not owner or comment not found",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

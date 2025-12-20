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


func abortWithError(c *gin.Context, statusCode int, errorCode, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error":   errorCode,
		"message": message,
	})
}

func getAuthContext(c *gin.Context) (model.AuthContext, bool) {
	authCtx, exists := c.Get("auth")
	if !exists {
		abortWithError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Authentication required")
		return model.AuthContext{}, false
	}
	return authCtx.(model.AuthContext), true
}

func validateEmailVerification(c *gin.Context, auth model.AuthContext, action string) bool {
	if !auth.EmailVerified {
		abortWithError(c, http.StatusForbidden, "EMAIL_NOT_VERIFIED",
			"Email verification required to "+action+" blogs")
		return false
	}
	return true
}

func parseUUIDParam(c *gin.Context, paramName string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(paramName))
	if err != nil {
		abortWithError(c, http.StatusBadRequest, "INVALID_ID", "Invalid blog ID")
		return uuid.Nil, false
	}
	return id, true
}

func bindJSON(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		abortWithError(c, http.StatusBadRequest, "BAD_REQUEST", "Invalid request body")
		return false
	}
	return true
}

func (h *BlogHandler) GetBlogs(c *gin.Context) {
	blogs, err := h.service.FindAll()
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to fetch blogs")
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) GetBlogByID(c *gin.Context) {
	blogID, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}

	blog, err := h.service.GetByID(blogID)
	if err != nil {
		abortWithError(c, http.StatusNotFound, "NOT_FOUND", "Blog not found")
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	auth, ok := getAuthContext(c)
	if !ok {
		return
	}

	if !validateEmailVerification(c, auth, "create") {
		return
	}

	var req model.CreateBlogRequest
	if !bindJSON(c, &req) {
		return
	}

	blog, err := h.service.CreateBlog(
		auth.UserID,
		req.Title,
		req.Content,
	)
	if err != nil {
		abortWithError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create blog")
		return
	}

	c.JSON(http.StatusCreated, blog)
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	auth, ok := getAuthContext(c)
	if !ok {
		return
	}

	if !validateEmailVerification(c, auth, "update") {
		return
	}

	blogID, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}

	var req model.UpdateBlogRequest
	if !bindJSON(c, &req) {
		return
	}

	updatedBlog, err := h.service.UpdateBlog(
		blogID,
		auth.UserID,
		req.Title,
		req.Content,
	)
	if err != nil {
		abortWithError(c, http.StatusForbidden, "FORBIDDEN", err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	auth, ok := getAuthContext(c)
	if !ok {
		return
	}

	if !validateEmailVerification(c, auth, "delete") {
		return
	}

	blogID, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}

	err := h.service.DeleteBlog(blogID, auth.UserID)
	if err != nil {
		abortWithError(c, http.StatusForbidden, "FORBIDDEN", err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

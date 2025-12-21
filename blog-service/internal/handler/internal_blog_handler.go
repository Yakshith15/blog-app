package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Yakshith15/blog-app/blog-service/internal/service"
)

type InternalBlogHandler struct {
	service *service.BlogService
}

func NewInternalBlogHandler(service *service.BlogService) *InternalBlogHandler {
	return &InternalBlogHandler{service: service}
}

func (h *InternalBlogHandler) CheckBlogExists(c *gin.Context) {

	blogID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	exists, err := h.service.BlogExists(blogID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !exists {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

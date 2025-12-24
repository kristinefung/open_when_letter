package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"open_when_letter/internal/service"
)

type LetterBoxHandler struct {
	service service.LetterBoxService
}

func NewLetterBoxHandler(service service.LetterBoxService) *LetterBoxHandler {
	return &LetterBoxHandler{service: service}
}

type CreateLetterBoxRequest struct {
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description,omitempty"`
	CoverImageURL string `json:"coverImageUrl,omitempty"`
}

func (h *LetterBoxHandler) CreateLetterBox(c *gin.Context) {
	var req CreateLetterBoxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	ownerID := "user-123-temp" // TODO: Replace with real user ID from JWT

	box, err := h.service.CreateLetterBox(req.Title, req.Description, req.CoverImageURL, ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create box"})
		return
	}

	c.JSON(http.StatusCreated, box)
}

package handlers

import (
	"codeDrop/internal/dto"
	"codeDrop/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PasteHandler struct {
	Service *service.PasteService
}

func NewAuthHandler (service *service.PasteService) *PasteHandler {
	return &PasteHandler{Service: service}
}

func (h *PasteHandler) CreatePaste(c *gin.Context) {
	var req dto.PasteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid credentials", })
		return
	}

	uId := c.GetString("user_id")
	userID, err := uuid.Parse(uId)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}

	paste, err := h.Service.Create(userID, req.Content, req.Visibility)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": paste.ID, "url": "/paste" + paste.ID, })
}

func (h *PasteHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	content, err := h.Service.FindById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "paste not found"})
		return
	}

	c.JSON(200, gin.H{"paste": content})
}

func (h *PasteHandler) GetByUser(c *gin.Context) {
	uid := c.GetString("user_id")
	userID, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id", })
		return
	}

	pastes, err := h.Service.FindByUser(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch paste", })
		return
	}
	c.JSON(200, pastes)
}

func (h *PasteHandler) DeletePaste(c *gin.Context) {
	id := c.Param("id")

	uid := c.GetString("user_id")
	userID, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(400, gin.H{"error": "user not found", })
		return
	}

	err = h.Service.DeletePaste(id, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "paste delted", })
}
package handlers

import (
	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateId(c *gin.Context) {
	if id, err := h.s.AddUser(); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	} else {
		c.JSON(200, gin.H{"id": id})
	}
}

func (h *Handlers) ViewInfo(c *gin.Context) {
	var id helpers.User
	var segment []string
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	if err := h.s.View(id, &segment); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"segment": segment})
}

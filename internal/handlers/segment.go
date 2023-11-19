package handlers

import (
	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateSeg(c *gin.Context) {
	var segment helpers.Segment
	if err := c.BindJSON(&segment); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := h.s.SegmentCr(segment); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
}

func (h *Handlers) DeleteSeg(c *gin.Context) {
	var segment helpers.Segment
	if err := c.ShouldBindUri(&segment); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := h.s.SegmentDel(segment); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
}

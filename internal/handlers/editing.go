package handlers

import (
	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) SegmentEd(c *gin.Context) {
	var add helpers.Add
	if err := c.BindJSON(&add); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := h.s.Addsegments(add); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := h.s.Delsegments(add); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
}

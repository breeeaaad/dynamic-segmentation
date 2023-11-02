package handlers

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Download(c *gin.Context) {
	data := [][]string{{"user_id", "segment_name", "procedure", "date"}}
	var datad [][]string
	date := c.Query("date")
	if err := h.s.GetAdd(date, &data); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	if err := h.s.GetDel(date, &datad); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	data = append(data, datad...)
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=report.csv")
	wr := csv.NewWriter(c.Writer)
	if err := wr.WriteAll(data); err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}
}

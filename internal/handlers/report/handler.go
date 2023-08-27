package report

import (
	"context"
	"encoding/csv"
	"fmt"

	config "github.com/breeeaaad/dynamic-segmentation/configs"
	"github.com/breeeaaad/dynamic-segmentation/internal/repository"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	data := [][]string{{"user_id", "segment_name", "procedure", "date"}}
	var datad [][]string
	date := c.Query("date")
	conn := config.Config()
	defer conn.Close(context.Background())
	if err := repository.GetAdd(conn, date, &data); err != nil {
		c.JSON(400, gin.H{"msg": err})
		fmt.Print(err)
		return
	}
	if err := repository.GetDel(conn, date, &datad); err != nil {
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

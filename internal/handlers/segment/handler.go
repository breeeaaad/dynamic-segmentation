package segment

import (
	"context"
	"fmt"

	config "github.com/breeeaaad/dynamic-segmentation/configs"
	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/breeeaaad/dynamic-segmentation/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateSeg(c *gin.Context) {
	var segment helpers.Segment
	if err := c.BindJSON(&segment); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	conn := config.Config()
	defer conn.Close(context.Background())
	if err := repository.SegmentCr(conn, segment); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
}

func DeleteSeg(c *gin.Context) {
	var segment helpers.Segment
	if err := c.ShouldBindUri(&segment); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	conn := config.Config()
	defer conn.Close(context.Background())
	if err := repository.SegmentDel(conn, segment); err != nil {
		c.JSON(400, gin.H{"msg": err})
		fmt.Print(err)
		return
	}
}

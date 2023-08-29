package editing

import (
	"context"
	"fmt"

	config "github.com/breeeaaad/dynamic-segmentation/configs"
	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/breeeaaad/dynamic-segmentation/internal/repository"
	"github.com/gin-gonic/gin"
)

func SegmentEd(c *gin.Context) {
	var add helpers.Add
	if err := c.BindJSON(&add); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	conn := config.Config()
	defer conn.Close(context.Background())
	if err := repository.Addsegments(conn, add); err != nil {
		c.JSON(400, gin.H{"msg": err})
		fmt.Print(err)
		return
	}
	if err := repository.Delsegments(conn, add); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
}

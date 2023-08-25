package account

import (
	"context"

	config "github.com/breeeaaad/dynamic-segmentation/configs"
	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/breeeaaad/dynamic-segmentation/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateId(c *gin.Context) {
	var id helpers.User
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	conn := config.Config()
	defer conn.Close(context.Background())
	if err := repository.AddUser(conn, id); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
}

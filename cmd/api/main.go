package main

import (
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/account"
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/editing"
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/segment"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/:id", account.CreateId)
	r.POST("/create", segment.CreateSeg)
	r.DELETE("/delete/:segment", segment.DeleteSeg)
	r.POST("/adding", editing.SegmentEd)
	r.Run()
}

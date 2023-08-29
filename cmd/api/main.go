package main

import (
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/account"
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/editing"
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/report"
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers/segment"
	"github.com/breeeaaad/dynamic-segmentation/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	go repository.Bg()
	r := gin.Default()
	r.POST("/newuser", account.CreateId)
	r.POST("/create", segment.CreateSeg)
	r.DELETE("/delete/:segment", segment.DeleteSeg)
	r.POST("/editing", editing.SegmentEd)
	r.GET("/:id", account.ViewInfo)
	r.GET("/download", report.Download)
	r.Run()
}

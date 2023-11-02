package main

import (
	"context"

	config "github.com/breeeaaad/dynamic-segmentation/configs"
	"github.com/breeeaaad/dynamic-segmentation/internal/handlers"
	"github.com/breeeaaad/dynamic-segmentation/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := config.Config()
	defer conn.Close(context.Background())
	repository := repository.New(conn)
	handlers := handlers.New(repository)
	go repository.Bg()
	r := gin.Default()
	r.POST("/newuser", handlers.CreateId)
	r.POST("/create", handlers.CreateSeg)
	r.DELETE("/delete/:segment", handlers.DeleteSeg)
	r.POST("/editing", handlers.SegmentEd)
	r.GET("/:id", handlers.ViewInfo)
	r.GET("/download", handlers.Download)
	r.Run()
}

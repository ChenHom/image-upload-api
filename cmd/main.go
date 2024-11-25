package main

import (
	"image-upload-api/internal/handlers"
	"image-upload-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化服務
	storageService := services.NewStorageService("./storage")
	aiProcessingService := services.NewAIProcessingService()

	// 創建 UploadHandler 的實例
	uploadHandler := &handlers.UploadHandler{
		StorageService:      *storageService,
		AIProcessingService: *aiProcessingService,
	}

	// 設定路由
	router := gin.New()
	router.POST("/upload", uploadHandler.HandleUpload)

	// 啟動服務器並監聽端口 8080
	router.Run(":8080")
}

package handlers

import (
	"context"
	"image-upload-api/internal/services"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

const maxUploadSize = 5 * 1024 * 1024 // 5 MB

type UploadHandler struct {
	StorageService      services.StorageService
	AIProcessingService services.AIProcessingService
}

func (h *UploadHandler) HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無法獲取檔案"})
		return
	}

	if file.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "檔案大小超過限制 (5 MB)"})
		return
	}

	// 安全性檢查邏輯（可根據需求實作）
	if err := h.checkFileSafety(file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "檔案不安全"})
		return
	}

	// 打開檔案
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法打開檔案"})
		return
	}
	defer src.Close()

	// 讀取檔案內容
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法讀取檔案"})
		return
	}

	// 儲存檔案
	err = h.StorageService.SaveToLocal(file.Filename, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "檔案儲存失敗"})
		return
	}

	// 異步處理圖檔
	go h.AIProcessingService.ProcessImage(context.Background(), file.Filename)

	c.JSON(http.StatusOK, gin.H{"message": "檔案上傳成功"})
}

func (h *UploadHandler) checkFileSafety(file *multipart.FileHeader) error {
	// 實作檔案安全性檢查邏輯
	return nil
}

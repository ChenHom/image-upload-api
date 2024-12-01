package handlers

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"image-upload-api/internal/services"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const maxUploadSize = 5 * 1024 * 1024 // 5 MB
var allowedExtensions = []string{".jpg", ".jpeg", ".png"}
var magicBytes = map[string][]byte{
	".jpg":  {0xFF, 0xD8, 0xFF},
	".jpeg": {0xFF, 0xD8, 0xFF},
	".png":  {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
}

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

	// 檢查副檔名
	if !h.isValidExtension(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支援的檔案類型"})
		return
	}

	// 防止目錄穿越
	if filepath.Base(file.Filename) != file.Filename {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法的檔案名"})
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

	// 驗證圖片內容
	if !h.isValidImage(fileBytes) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法的圖片內容"})
		return
	}

	// 驗證 magic bytes
	if !h.isValidMagicBytes(fileBytes, file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法的圖片內容"})
		return
	}

	// 計算檔案指紋
	hash := sha256.Sum256(fileBytes)
	fileFingerprint := hex.EncodeToString(hash[:])
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// 儲存檔案
	err = h.StorageService.SaveToLocal(fileFingerprint, ext, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "檔案儲存失敗"})
		return
	}

	// 異步處理圖檔
	go h.AIProcessingService.ProcessImage(context.Background(), fileFingerprint)

	c.JSON(http.StatusOK, gin.H{"message": "檔案上傳成功"})
}

func (h *UploadHandler) checkFileSafety(file *multipart.FileHeader) error {
	// 實作檔案安全性檢查邏輯
	return nil
}

func (h *UploadHandler) isValidExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

func (h *UploadHandler) isValidImage(fileBytes []byte) bool {
	_, _, err := image.DecodeConfig(bytes.NewReader(fileBytes))
	return err == nil
}

func (h *UploadHandler) isValidMagicBytes(fileBytes []byte, filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	magic, exists := magicBytes[ext]
	if !exists {
		return false
	}
	return bytes.HasPrefix(fileBytes, magic)
}

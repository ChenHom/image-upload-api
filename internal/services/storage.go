package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type StorageService struct {
	storagePath string
}

func NewStorageService(storagePath string) *StorageService {
	return &StorageService{storagePath: storagePath}
}

func (s *StorageService) SaveToLocal(fileFingerprint, ext string, data []byte) error {
	if len(data) > 5*1024*1024 { // 檔案大小檢查，限制最大 5 MB
		return errors.New("檔案大小超過限制")
	}

	// 確認資料夾是否存在，不存在則新增
	if _, err := os.Stat(s.storagePath); os.IsNotExist(err) {
		if err := os.MkdirAll(s.storagePath, os.ModePerm); err != nil {
			return errors.New("無法建立資料夾")
		}
	}

	// 保留原本的副檔名
	filePath := filepath.Join(s.storagePath, fileFingerprint+ext)
	fmt.Printf("儲存 %s\n", fileFingerprint+ext)
	fmt.Printf("儲存檔案到 %s\n", filePath)
	return os.WriteFile(filePath, data, 0644)
}

func (s *StorageService) SaveToCloud(fileName string, data []byte) error {
	// 實作雲端儲存邏輯
	return errors.New("雲端儲存尚未實作")
}

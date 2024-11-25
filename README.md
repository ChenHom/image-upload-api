# Image Upload API

這是一個使用 Go 語言建立的圖檔上傳 API，並使用 Nginx 作為 Web 伺服器。該專案提供了一個簡單的介面，允許用戶上傳圖檔，並進行檔案大小和安全性檢查。

## 專案結構

```
image-upload-api
├── cmd
│   └── main.go          # 應用程式的進入點，負責啟動 HTTP 伺服器
├── internal
│   ├── handlers
│   │   └── upload.go    # 上傳處理邏輯
│   ├── services
│   │   ├── storage.go    # 檔案儲存服務
│   │   └── ai_processing.go # AI 圖像處理服務
├── configs
│   └── config.go        # 應用程式配置設定
├── nginx.conf           # Nginx 配置檔
├── go.mod               # Go 模組配置檔
├── go.sum               # Go 模組依賴項校驗和
└── README.md            # 專案文檔
```

## 安裝

1. 確保已安裝 Go 語言環境。
2. 下載專案檔案：
   ```
   git clone <repository-url>
   cd image-upload-api
   ```
3. 安裝依賴項：
   ```
   go mod tidy
   ```

## 配置

在 `configs/config.go` 中設定應用程式的配置選項，包括儲存路徑和雲端服務的相關設定。

## 使用

1. 啟動應用程式：
   ```
   go run cmd/main.go
   ```
2. 使用 HTTP 客戶端（如 Postman）向 `/upload` 路徑發送 POST 請求，並附上圖檔。

## 注意事項

- 上傳的圖檔大小不得超過 5 MB。
- 圖檔將進行安全性檢查，確保符合上傳規則。
- 預設情況下，檔案將儲存到本地，若需儲存至雲端，請在配置中進行相應設定。

## 貢獻

歡迎任何形式的貢獻，請提出問題或提交拉取請求。
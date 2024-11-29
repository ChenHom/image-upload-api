# 第一階段：編譯程式碼
FROM golang:1.23-alpine AS builder

# 設置工作目錄
WORKDIR /app

# 複製模組管理文件並安裝依賴
COPY go.mod go.sum ./
RUN go mod download

# 設置環境變數
ARG MODE=production

# 複製所有程式碼
COPY . ./

# 僅在 production 模式下編譯 Go 程式
RUN if [ "$MODE" = "production" ]; then \
      CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go; \
    fi

# 第二階段：運行環境
ARG MODE=production
FROM golang:1.23-alpine

# 設置工作目錄
WORKDIR /root/

# 在 debug 模式下從 builder 階段拷貝源代碼
COPY --from=builder /app /app

# 僅在 production 模式下從 builder 階段拷貝編譯好的二進位檔案
RUN if [ "$MODE" = "production" ]; then \
      cp /app/main .; \
    fi

# 複製啟動腳本
COPY entrypoint.sh .

# 暴露埠（如果需要）
EXPOSE 8080

# 設置啟動腳本的執行權限
RUN chmod +x /root/entrypoint.sh

# 啟動應用
CMD ["/root/entrypoint.sh"]
# 第一階段：編譯程式碼
FROM golang:1.23-alpine AS builder

# 設置工作目錄
WORKDIR /app

# 複製模組管理文件並安裝依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製所有程式碼
COPY . ./

# 編譯 Go 程式
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# 第二階段：運行環境
FROM alpine:latest

# 設置工作目錄
WORKDIR /root/

# 從 builder 階段拷貝編譯好的二進位檔案
COPY --from=builder /app/main .

# 暴露埠（如果需要）
EXPOSE 8080

# 啟動應用
CMD ["/root/main"]
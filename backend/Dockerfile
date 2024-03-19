# ベースイメージを指定
# ビルド環境
FROM golang:1.18 as builder
WORKDIR /app

# go.mod と go.sum をコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係を解決
RUN go mod tidy

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# 実行環境
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]

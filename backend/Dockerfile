# ビルドステージ
FROM golang:1.18 AS builder
WORKDIR /app

# モジュールのコピー
COPY go.mod .
COPY go.sum .

# 依存関係のダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

COPY .env ./

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main .

# 実行ステージ
FROM alpine:latest
WORKDIR /root/

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .

# ポート8080を公開
EXPOSE 8080

# 実行コマンド
CMD ["./main"]


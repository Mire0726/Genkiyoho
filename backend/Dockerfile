# ベースイメージを指定
FROM golang:latest

LABEL maintainer="SumireInoue <violeta.jin.726@gmail.com>"
# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# アプリケーションをビルド
RUN go build -o main .

# コンテナのポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD [ "./main"]

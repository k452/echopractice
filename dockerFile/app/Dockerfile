# goとalpine
FROM golang:1.13.6-alpine

# update
RUN apk add --update && apk add git

# TimeZoneをJSTに変更
ENV TZ=Asia/Tokyo
RUN apk --no-cache add tzdata

# goの作業ディレクトリ作成
RUN mkdir /go/src/app

# 作業ディレクトリ指定 
WORKDIR /go/src/app

# ローカルのファイルをコンテナへ移動
ADD . /go/src/app

# golangの必要ライブラリをインストール
RUN go get github.com/labstack/echo/... && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/sirupsen/logrus && \
    go get github.com/joho/godotenv && \
    GO111MODULE=off go get github.com/oxequa/realize

EXPOSE 8080
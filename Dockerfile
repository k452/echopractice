# goとalpine
FROM golang:1.13.6-alpine

# update
RUN apk add --update && apk add git

# goの作業ディレクトリ作成
RUN mkdir /go/src/app

# 作業ディレクトリ指定 
WORKDIR /go/src/app

# ローカルのファイルをコンテナへ移動
ADD . /go/src/app

# golangの必要ライブラリをインストール
RUN go get -u github.com/labstack/echo/...

EXPOSE 8080
# さけ.io の API

## 使用技術

- Go
  - Echo
  - realize
  - logrus
  - godotenv
- MySQL

## 開発環境

- Docker
- Docker-Compose

## ディレクトリ構成

├── README.md  
├── conf --- 環境変数の読み込み  
│   └── dbConfig.go  
├── db --- db 接続  
│   └── connect.go  
├── docker-compose.yml  
├── Dockerfile  
├── envFile --- 環境変数ファイル  
│   ├── production.env  
│   └── test.env  
├── go.mod  
├── go.sum  
├── main.go --- API サーバー起動  
├── model  
│   ├── init.go --- テーブル作成等の初期化処理  
│   ├── management.go  
│   ├── recipe.go  
│   ├── type.go  
│   └── users.go  
├── mysql_data  
│   └── 省略  
├── routing --- ルーティング  
│   ├── main.go  
│   ├── management.go  
│   ├── recipe.go  
│   └── users.go  
└── sqls  
   └── 省略

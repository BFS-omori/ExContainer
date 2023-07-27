# はじめに

## 概要

Window10 + WSL + DockerDesktop環境でのコンテナの実演デモ用のプロジェクト

## ディレクトリ構成

```text
.
├── build       Dockerfile他
├── deploy      docker-compose.yml
└── src         ソースコード
    ├── go1     go バッチ
    └── go2     go API
```

## コマンド

- Goコマンド

```sh
#Go 実行
go run ./src/go1/main.go
#Go ビルド
go build -o ./bin/go1/ ./src/go1/main.go
#実行
bin/go1/main
#Go ビルド
go build -o ./bin/go2/httpd ./src/go2/main.go
```

- Dockerコマンド1(go1バッチ image名 exc_go1 コンテナ名 exc_go1a)

```sh
#Dockerイメージビルド
docker image build -t exc_go1 -f ./build/go1/Dockerfile ./
#コンテナ実行（作成&開始）
docker container run --name exc_go1a -d exc_go1
#ログ確認
docker logs exc_go1a -f
#コンテナに入る
docker exec -it exc_go1a sh
#コンテナ削除
docker rm exc_go1a

#Dockerイメージ確認
docker images | grep exc
#コンテナ確認
docker ps | grep exc
#コンテナ作成
docker create --name exc_go1a -t -i exc_go1
#コンテナ開始
docker start -a -i exc_go1a
#コンテナ保存
docker save exc_go1 > ./images/exc_go1.tar
#ファイルサイズ確認
ls -lh images/
```

- Dockerコマンド2(go2API image名 exc_go2 コンテナ名 exc_go2a)

```sh
#Dockerイメージビルド
docker image build -t exc_go2 -f ./build/go2/Dockerfile ./
#コンテナ実行（作成&開始）Port3000を指定して実行
docker container run --name exc_go2a -p 3000:8080 -d exc_go2
#コンテナ確認
docker ps | grep go2
#CURLでJSONリクエスト
curl -X POST -H "Content-Type: application/json" -d '{"ID": "AAA", "PWD": "BBB"}' http://localhost:3000/Auth | jq .
#ログ確認
docker logs exc_go2a -f
#コンテナに入る
docker exec -it exc_go2a sh
#コンテナ削除
docker rm exc_go2a
```

- docker-compose.ymlを使用した起動

```sh
cd deploy
#コンテナUP
dcs -f docker-compose.yml up -d
dcs -f docker-compose.yml down
#コンテナ一覧
docker ps
#CURLでJSONリクエスト
curl -X POST -H "Content-Type: application/json" -d '{"ID": "AAA", "PWD": "BBB"}' http://localhost:18010/Auth | jq .
#ログ
docker logs exc3
#統計情報（CPU,メモリ利用/制限,メモリ,NET,ブロックデバイス）
docker stats
```

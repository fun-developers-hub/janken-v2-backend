# janken-v2-backend

## 必要要件

- Go 1.26.0 以上

## セットアップ

```bash
# 依存関係をインストール
go mod download

# 環境変数ファイルを用意
cp .env.example .env
# .env を編集して値を設定する
```

## 環境変数

| 変数名 | デフォルト値 | 説明 |
| --- | --- | --- |
| `PORT` | `8080` | サーバーがリッスンするポート番号 |
| `ALLOW_ORIGINS` | `*` | CORSで許可するオリジン（カンマ区切りで複数指定可能） |

## 起動

```bash
go run cmd/server/main.go
```

起動後、`http://localhost:8080` でサーバーにアクセスできます。

## ビルド

```bash
go build -o bin/server cmd/server/main.go
# bin/server にバイナリが生成される
```

## Swagger UI

```bash
# docs/ 以下のAPI定義を再生成する場合
make swag
```

起動中のサーバーの `/swagger` にアクセスすると Swagger UI が表示されます。

## その他

```bash
go mod tidy
go fmt ./...
```

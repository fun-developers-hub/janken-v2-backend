# janken-v2-backend

じゃんけんゲームのバックエンド API。

## 必要要件

- Go 1.26 以上
- Docker / Docker Compose（MySQL を含めて起動する場合）

## 構成

```
cmd/server/main.go     # エントリ + DI 配線
internal/
  config/              # 環境変数・DSN
  handler/             # echo ハンドラ（swag アノテーション）
  infra/mysql/         # DB 接続
  server/              # echo 初期化・ルーティング
migrations/            # golang-migrate の SQL
```

ドメインを足すときは `internal/domain`（エンティティ + repository interface）と `internal/usecase` を追加し、`handler → usecase → domain ← infra` の向きで配線する。

## セットアップ・起動

```bash
cp .env.example .env       # 値を編集
make up                    # Go + MySQL を起動
make migrate               # マイグレーション適用（別シェル）
```

- 確認: `curl localhost:8080/health` / `curl localhost:8080/health/db`
- Swagger UI: http://localhost:8080/swagger/index.html
- 停止: `make down`（データ保持） / `make down-v`（データ削除） / ログ: `make logs`

## 環境変数

| 変数 | デフォルト | 説明 |
| --- | --- | --- |
| `PORT` | `8080` | リッスンポート |
| `ALLOW_ORIGINS` | `*` | CORS 許可オリジン |
| `DB_HOST` / `DB_PORT` | `127.0.0.1` / `3306` | MySQL 接続先（compose では host=`mysql`） |
| `DB_USER` / `DB_PASSWORD` / `DB_NAME` | `app` / `password` / `janken` | MySQL 認証・DB 名 |
| `DB_URL` | — | migrate 用 URL（`mysql://user:pass@tcp(host:port)/db`） |

## マイグレーション

DB のスキーマは `migrations/` の SQL で管理する。基本は **`make migrate` を叩くだけ**。

```bash
make migrate                         # 最新まで適用（DBの準備はこれだけでOK）
make migrate-create name=create_xxx  # テーブル追加用の SQL ファイルを作る
```

新しいテーブルを足すときは `make migrate-create name=...` で `migrations/000002_xxx.up.sql`（と `.down.sql`）が生成されるので、`up.sql` に `CREATE TABLE ...` を書いて `make migrate` する。

`migrations/000001_create_sample.*` は動作確認用のサンプル。差し替え/削除してよい。

## 開発

```bash
make swag          # OpenAPI 定義を再生成
go test ./...
```

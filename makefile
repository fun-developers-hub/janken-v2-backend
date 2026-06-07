SWAG       := go run github.com/swaggo/swag/v2/cmd/swag@latest
SWAG_ENTRY := cmd/server/main.go

# golang-migrate も swag と同様 go run で実行(別途インストール不要)
MIGRATE     := go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
MIGRATE_DIR := migrations
DB_HOST     ?= 127.0.0.1
DB_PORT     ?= 3306
DB_USER     ?= app
DB_PASSWORD ?= password
DB_NAME     ?= janken
DB_URL      ?= mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

.PHONY: swag
swag: ## API定義(OpenAPI 3.1)を docs/ に生成し /swagger で配信
	$(SWAG) init -g $(SWAG_ENTRY) --v3.1

.PHONY: up
up: swag ## Go と MySQL を docker compose で起動 (先に swag で docs/ を生成)
	docker compose up --build

.PHONY: down
down: ## docker compose を停止 (MySQLデータは volume に保持)
	docker compose down

.PHONY: down-v
down-v: ## docker compose を停止し MySQLデータ(volume)も削除
	docker compose down -v

.PHONY: logs
logs: ## コンテナのログを追跡表示
	docker compose logs -f

.PHONY: migrate
migrate: ## マイグレーションを最新まで適用（DBの準備はこれだけでOK）
	$(MIGRATE) -path $(MIGRATE_DIR) -database "$(DB_URL)" up

.PHONY: migrate-create
migrate-create: ## 新しいマイグレーションファイルを作る 例: make migrate-create name=create_players
	$(MIGRATE) create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

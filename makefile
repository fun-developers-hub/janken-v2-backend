SWAG       := go run github.com/swaggo/swag/v2/cmd/swag@latest
SWAG_ENTRY := cmd/server/main.go

.PHONY: swag
swag: ## API定義(OpenAPI 3.1)を docs/ に生成し /swagger で配信
	$(SWAG) init -g $(SWAG_ENTRY) --v3.1

# ---- build stage ----
FROM golang:1.26 AS builder

WORKDIR /src

# 依存だけ先に取得してレイヤキャッシュを効かせる
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# docs/ は .gitignore 対象なのでビルド内で生成する(make 不要で Docker 単体ビルド可能)
RUN go run github.com/swaggo/swag/v2/cmd/swag@latest init -g cmd/server/main.go --v3.1
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /bin/server ./cmd/server

# ---- run stage ----
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /bin/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]

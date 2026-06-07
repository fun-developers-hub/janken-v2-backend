package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fun-developers-hub/janken-v2-backend/internal/config"
	"github.com/fun-developers-hub/janken-v2-backend/internal/handler"
	"github.com/fun-developers-hub/janken-v2-backend/internal/infra/mysql"
	"github.com/fun-developers-hub/janken-v2-backend/internal/server"
)

// @title       Janken Backend API
// @version     0.1.0
// @description じゃんけんゲームのバックエンドAPI
// @servers.url http://localhost:8080
func main() {
	// 本番コンテナ(distroless)の healthcheck 用セルフチェック。
	if len(os.Args) > 1 && os.Args[1] == "-healthcheck" {
		healthcheck()
		return
	}

	cfg := config.Load()

	database, err := mysql.Open(cfg.DB)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer database.Close()

	// --- DI 配線(ここだけが全層を組み立てる) ---
	s := server.New(cfg, server.Handlers{
		Health: handler.NewHealthHandler(database),
	})
	log.Fatal(s.Serve(cfg.Port))
}

// healthcheck は自身の /health を叩き、200 以外なら非ゼロ終了する。
// distroless イメージには curl/wget が無いため、バイナリ自身で代替する。
func healthcheck() {
	port := config.Load().Port
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://127.0.0.1:" + port + "/health")
	if err != nil || resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	os.Exit(0)
}

package main

import (
	"log"

	"github.com/fun-developers-hub/janken-v2-backend/app"
	"github.com/fun-developers-hub/janken-v2-backend/app/config"
)

// @title       Janken Backend API
// @version     0.1.0
// @description じゃんけんゲームのバックエンドAPI
// @servers.url http://localhost:8080
func main() {
	cfg := config.Load()
	s := app.NewServer(cfg)
	log.Fatal(s.Serve(cfg.Port))
}

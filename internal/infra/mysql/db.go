package mysql

import (
	"database/sql"
	"time"

	"github.com/fun-developers-hub/janken-v2-backend/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Open(cfg config.DBConfig) (*sql.DB, error) {
	d, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}
	d.SetConnMaxLifetime(3 * time.Minute)
	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(10)
	return d, nil
}

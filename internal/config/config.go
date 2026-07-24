package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port         string
	AllowOrigins []string
	DB           DBConfig
	PlayLogPath  string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name,
	)
}

func Load() Config {
	return Config{
		Port:         getenv("PORT", "8080"),
		AllowOrigins: splitOrDefault(os.Getenv("ALLOW_ORIGINS"), []string{"*"}),
		DB: DBConfig{
			Host:     getenv("DB_HOST", "127.0.0.1"),
			Port:     getenv("DB_PORT", "3306"),
			User:     getenv("DB_USER", "app"),
			Password: getenv("DB_PASSWORD", "password"),
			Name:     getenv("DB_NAME", "janken"),
		},
		PlayLogPath: "/app-log/play_log.csv",
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func splitOrDefault(v string, def []string) []string {
	if v == "" {
		return def
	}
	return strings.Split(v, ",")
}

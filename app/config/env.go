package config

import (
	"os"
	"strings"
)

type Config struct {
	Port         string
	AllowOrigins []string
}

func Load() Config {
	return Config{
		Port:         getenv("PORT", "8080"),
		AllowOrigins: splitOrDefault(os.Getenv("ALLOW_ORIGINS"), []string{"*"}),
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

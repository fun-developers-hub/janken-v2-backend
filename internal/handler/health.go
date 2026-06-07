package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

type HealthResponse struct {
	Status string `json:"status"`
}

// Health godoc
// @Summary  ヘルスチェック
// @Tags     System
// @Produce  json
// @Success  200 {object} HealthResponse
// @Router   /health [get]
func (h *HealthHandler) Health(c *echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status: "ok",
	})
}

// DBHealth godoc
// @Summary  DBヘルスチェック (MySQLへの疎通確認)
// @Tags     System
// @Produce  json
// @Success  200 {object} HealthResponse
// @Failure  503 {object} HealthResponse
// @Router   /health/db [get]
func (h *HealthHandler) DBHealth(c *echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	if err := h.db.PingContext(ctx); err != nil {
		return c.JSON(http.StatusServiceUnavailable, HealthResponse{Status: "error"})
	}
	return c.JSON(http.StatusOK, HealthResponse{Status: "ok"})
}

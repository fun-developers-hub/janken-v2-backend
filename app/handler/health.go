package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthHandler struct{}

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

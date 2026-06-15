package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type JankenCounterHandler struct {
}

type JankenCounterResponse struct {
	Count int `json:"count"`
}

func NewJankenCounterHandler() *JankenCounterHandler {
	return &JankenCounterHandler{}
}

// JankenCounter godoc
// @Summary  じゃんけんの回数を出力する
// @Tags     System
// @Produce  json
// @Success  200 {object} JankenCounterResponse
// @Router   /janken/counter [get]
func (h *JankenCounterHandler) JankenCounter(c *echo.Context) error {
	return c.JSON(http.StatusOK, JankenCounterResponse{
		Count: 1,
	})
}

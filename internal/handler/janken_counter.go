package handler

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v5"
)

type GameSession struct {
	mu    sync.Mutex
	count int
}

func (s *GameSession) Current() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.count
}

func (s *GameSession) Advance() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.count == 5 {
		s.count = 0
	}

	s.count++
}

type JankenCounterHandler struct {
	session *GameSession
}

type JankenCounterResponse struct {
	Count int `json:"count"`
}

func NewJankenCounterHandler(session *GameSession) *JankenCounterHandler {
	return &JankenCounterHandler{
		session: session,
	}
}

// JankenCounter godoc
// @Summary  じゃんけんの回数を出力する
// @Tags     Janken
// @Produce  json
// @Success  200 {object} JankenCounterResponse
// @Router   /janken/counter [get]
func (h *JankenCounterHandler) JankenCounter(c *echo.Context) error {
	return c.JSON(http.StatusOK, JankenCounterResponse{
		Count: h.session.Current(),
	})
}

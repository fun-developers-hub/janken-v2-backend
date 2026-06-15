package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

// リクエストボディの構造体
type PlayGameRequest struct {
	UserHand string `json:"user_hand"` // "rock", "scissors", "paper"
}

// 200 OK: 成功時のレスポンス
type PlayGameResponse struct {
	Result  string `json:"result"`   // "win", "lose", "draw"
	CPUHand string `json:"cpu_hand"` // "rock", "scissors", "paper"
}

// PlayGame godoc
// @Summary     じゃんけんを実行する
// @Description ユーザーの手を受け取って、CPUとのじゃんけん対戦の結果を返してくれるAPI
// @Accept      json
// @Produce     json
// @Param       request body PlayGameRequest true "ユーザーの手"
// @Success     200 {object} PlayGameResponse
// @Failure     400 {object} map[string]string
// @Router      /janken [post]
func (g *GameHandler) PlayGame(c *echo.Context) error {
	// ユーザーの手を読み込む
	var req PlayGameRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rock, scissors, paperのいずれかを送信してください"})
	}

	// キーは任意の手、値はそれに負ける手
	winningMap := map[string]string{
		"rock":     "scissors",
		"scissors": "paper",
		"paper":    "rock",
	}

	// rock, scissors, paper以外を弾く
	if _, exists := winningMap[req.UserHand]; !exists {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rock, scissors, paperのいずれかを送信してください"})
	}

	// CPUの手をランダムに決める
	hands := []string{"rock", "scissors", "paper"}
	cpuHand := hands[rand.Intn(3)]

	// 勝敗を判定する
	result := "lose"

	if req.UserHand == cpuHand {
		result = "draw"
	} else if winningMap[req.UserHand] == cpuHand {
		result = "win"
	}

	// JSONで結果を返す
	res := PlayGameResponse{
		Result:  result,
		CPUHand: cpuHand,
	}
	return c.JSON(http.StatusOK, res)
}

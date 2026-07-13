package handler

import (
	"database/sql"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v5"
)

var winningMap = map[string]string{
	"rock":     "scissors",
	"scissors": "paper",
	"paper":    "rock",
}

type GameHandler struct {
	db *sql.DB
}

func NewGameHandler(db *sql.DB) *GameHandler {
	return &GameHandler{db: db}
}

// リクエストボディの構造体
type PlayGameRequest struct {
	UserHand string `json:"user_hand" enums:"rock,scissors,paper" example:"rock"` // "rock", "scissors", "paper"
}

// 200 OK: 成功時のレスポンス
type PlayGameResponse struct {
	Result  string `json:"result" enums:"win,lose,draw" example:"win"`              // "win", "lose", "draw"
	CPUHand string `json:"cpu_hand" enums:"rock,scissors,paper" example:"scissors"` // "rock", "scissors", "paper"
}

// 400 Bad Request: エラー時のレスポンス
type ErrorResponse struct {
	Error string `json:"error"` // エラーメッセージ
}

// PlayGame godoc
// @Summary     じゃんけんを実行する
// @Description ユーザーの手を受け取って、CPUとのじゃんけん対戦の結果を返してくれるAPI
// @Tags        Janken
// @Produce     json
// @Param       request body PlayGameRequest true "ユーザーの手"
// @Success     200 {object} PlayGameResponse
// @Failure     400 {object} ErrorResponse
// @Router      /janken [post]
func (g *GameHandler) PlayGame(c *echo.Context) error {
	// ユーザーの手を読み込む
	var req PlayGameRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "リクエストが不正です"})
	}

	// rock, scissors, paper以外を弾く
	if _, exists := winningMap[req.UserHand]; !exists {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "rock, scissors, paperのいずれかを送信してください"})
	}

	// CPUの手をランダムに決める
	hands := []string{"rock", "scissors", "paper"}
	cpuHand := hands[rand.Intn(3)]

	// 勝敗を判定する
	result := getMatchResult(req.UserHand, cpuHand)

	// JSONで結果を返す
	res := PlayGameResponse{
		Result:  result,
		CPUHand: cpuHand,
	}
	return c.JSON(http.StatusOK, res)
}

func getMatchResult(userHand, cpuHand string) string {
	ans := "lose"

	if userHand == cpuHand {
		ans = "draw"
	} else if winningMap[userHand] == cpuHand {
		ans = "win"
	}
	return ans
}

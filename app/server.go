package app

import (
	"github.com/fun-developers-hub/janken-v2-backend/app/config"
	"github.com/fun-developers-hub/janken-v2-backend/app/handler"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	echoSwagger "github.com/swaggo/echo-swagger/v2"

	// swag v2 が生成した OpenAPI 定義を登録する (init の副作用が目的。消すと /swagger が 500 になる)
	_ "github.com/fun-developers-hub/janken-v2-backend/docs"
)

type Server struct {
	e *echo.Echo
}

func NewServer(cfg config.Config) *Server {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS(cfg.AllowOrigins...))

	registerRoutes(e)

	return &Server{e}
}

func registerRoutes(e *echo.Echo) {
	health := &handler.HealthHandler{}
	e.GET("/health", health.Health)

	// OpenAPI 3.x (swag v2) 用の Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandlerV3)
}

func (s *Server) Serve(port string) error {
	return s.e.Start(":" + port)
}

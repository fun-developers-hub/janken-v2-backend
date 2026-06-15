package server

import (
	"github.com/fun-developers-hub/janken-v2-backend/internal/config"
	"github.com/fun-developers-hub/janken-v2-backend/internal/handler"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	echoSwagger "github.com/swaggo/echo-swagger/v2"

	_ "github.com/fun-developers-hub/janken-v2-backend/docs"
)

type Handlers struct {
	Health        *handler.HealthHandler
	JankenCounter *handler.JankenCounterHandler
}

type Server struct {
	e *echo.Echo
}

func New(cfg config.Config, h Handlers) *Server {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS(cfg.AllowOrigins...))

	registerRoutes(e, h)

	return &Server{e}
}

func registerRoutes(e *echo.Echo, h Handlers) {
	e.GET("/health", h.Health.Health)
	e.GET("/health/db", h.Health.DBHealth)
	e.GET("/janken/counter", h.JankenCounter.JankenCounter)

	// OpenAPI 3.x (swag v2) 用の Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandlerV3)
}

func (s *Server) Serve(port string) error {
	return s.e.Start(":" + port)
}

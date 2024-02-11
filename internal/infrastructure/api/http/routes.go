package http

import (
	"net/http"

	"github.com/i4erkasov/go-ddd-cqrs/internal/infrastructure/api/http/handler"
	"github.com/i4erkasov/go-ddd-cqrs/internal/infrastructure/api/http/middleware"
	"github.com/labstack/echo/v4"
)

func (s *Server) routes(h *handler.HttpHandler, m *middleware.HttpMiddleware) {
	s.echo.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	api := s.echo.Group("/api")
	api.GET("/hello", h.HelloWold)
}

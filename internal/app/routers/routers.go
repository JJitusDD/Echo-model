package routers

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"echo-model/internal/domain/service"
)

func Setup(e *echo.Echo, s *service.Service) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}

package routers

import (
	"net/http"

	"echo-model/internal/app/routers/user"
	"echo-model/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, s *service.Service) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	user.NewUserRouter(e, s)

}

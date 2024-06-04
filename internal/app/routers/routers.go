package routers

import (
	"echo-model/internal/app/routers/user"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"

	"echo-model/internal/domain/service"
)

func Setup(e *echo.Echo, s *service.Service) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.Config.JwtSecret),
	}))

	user.NewUserRouter(e, s)

}

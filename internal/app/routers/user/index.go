package user

import (
	"echo-model/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func NewRouter(g *echo.Echo, s *service.Service) {
	routes := g.Group("/users")

	// Login and logout with JWT
	routes.POST("/logout", s.UserLogout, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/login", s.UserLogin, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/refresh-token", s.UserRefreshToken, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/profile", s.UserProfile, s.ValidateAccesstoken(), s.VerifyHash())
}

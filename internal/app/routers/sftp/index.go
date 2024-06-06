package sftp

import (
	"echo-model/internal/domain/service"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(g *echo.Echo, s *service.Service) {
	routes := g.Group("/sftp")

	// Login and logout with JWT
	routes.POST("/read-path", s.SftpReadPath)
	routes.POST("/login", s.UserLogin, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/refresh-token", s.UserRefreshToken, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/profile", s.UserProfile, s.ValidateAccesstoken(), s.VerifyHash())
}

package user

import (
	"echo-model/internal/domain/service"
	"echo-model/pkg/helper/crypt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewUserRouter(g *echo.Echo, s *service.Service) {
	routes := g.Group("/users")

	// Login and logout with JWT
	jwt_config := middleware.JWTConfig{
		Claims:     &crypt.JwtCustomClaims{},
		SigningKey: []byte(s.Config.JwtSecret),
	}
	routes.POST("/logout", s.UserLogout, echojwt.JWT(jwt_config), s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/login", s.UserLogin, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/refresh-token", s.UserRefreshToken, s.ValidateAccesstoken(), s.VerifyHash())
	routes.POST("/profile", s.UserProfile, echojwt.JWT(jwt_config), s.ValidateAccesstoken(), s.VerifyHash())
}

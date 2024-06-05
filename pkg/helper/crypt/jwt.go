package crypt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type UserToken struct {
	UserId     string `json:"user_id"`
	DeviceId   string `json:"device_id"`
	Secret     string `json:"secret"`
	AppVersion string `json:"app_version"`
}

type JwtCustomClaims struct {
	Payload string `json:"payload"`
	jwt.Claims
}

// EncryptAccessToken returns a complete, signed JWT.
func EncryptAccessToken(payload, secret string) string {
	tokenTime := cast.ToDuration(5)
	if os.Getenv("token_time") != "" {
		tokenTime = cast.ToDuration(os.Getenv("token_time"))
	}

	claims := &JwtCustomClaims{
		Payload: payload,
		Claims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * tokenTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err.Error())
	}

	return t
}

func DecryptAccessToken(c echo.Context) (*JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)
	claims, ok := user.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("DecryptAccessToken Fail")
	}
	return claims, nil
}

func DecryptTokenByString(token string) (*JwtCustomClaims, error) {
	claims := &JwtCustomClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return claims, nil
}

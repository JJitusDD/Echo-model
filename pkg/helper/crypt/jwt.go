package crypt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"

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
	jwt.StandardClaims
}

func EncryptAccessToken(payload, secret string) string {
	coeffection := cast.ToDuration(5)
	if os.Getenv("Coeffection") != "" {
		coeffection = cast.ToDuration(os.Getenv("Coeffection"))
	}

	claims := &JwtCustomClaims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * coeffection).Unix(),
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

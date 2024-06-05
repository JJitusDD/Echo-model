package app

import (
	"strings"

	"echo-model/config"
	_ "echo-model/docs"
	"echo-model/internal/app/middleware"
	"echo-model/internal/app/routers"
	"echo-model/internal/domain/service"
	"echo-model/internal/infrastructure/facade"
	error_internal "echo-model/pkg/error"
	"echo-model/pkg/helper/crypt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/echo-swagger"
)

type App struct {
	Config *configs.Config
	Echo   *echo.Echo
	Facade *facade.EchoModelFacade
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func Initialize() (*App, error) {
	e := echo.New()
	confg, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	f := facade.NewEchoModelFacade(confg)
	e.Logger.SetLevel(log.INFO)

	if confg.ENV != "PRODUCTION" {
		f.Logger.WithFields(logrus.Fields{
			"config": confg,
		}).Info("Logging config")
	}

	// setup Validator
	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}

	// Middleware
	e.Use(echo_middleware.Recover())
	e.Use(echo_middleware.RequestID())
	e.Use(middleware.TraceIDMiddleware())
	e.Use(middleware.LogCollect(f.Logger))
	e.Use(echo_middleware.GzipWithConfig(echo_middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics") // Change "metrics" for your own path
		},
	}))

	// error handler
	e.HTTPErrorHandler = error_internal.CustomHTTPErrorHandler(f.Logger)

	newServ := service.NewService(f)

	// Custom jwt config
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(newServ.Config.JwtSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &crypt.JwtCustomClaims{}
		},
	}))

	routers.Setup(e, newServ)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &App{
		Config: confg,
		Echo:   e,
		Facade: f,
	}, nil
}

func (a *App) Run() {
	if err := a.Echo.Start(":" + a.Config.Port); err != nil {
		a.Echo.Logger.Info("shutting down the server")
	}
}

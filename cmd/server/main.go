package main

import (
	"echo-model/internal/app"
	"echo-model/pkg/wrapper"
)

// @title Swagger Echo model API
// @version 1.0
// @description This is a docs API for Echo model.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host echo-model-test
// @BasePath /
func main() {
	app, err := app.Initialize()
	if err != nil {
		panic(err)
	}
	go app.Run()
	// Graceful
	wrapper.GracefulShutdown(app.Echo)
}

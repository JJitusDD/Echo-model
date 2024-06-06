package utilties

import (
	"github.com/labstack/echo/v4"
)

func BindingBody(body interface{}, ctx echo.Context) error {
	if err := ctx.Bind(body); err != nil {
		return err
	}

	if err := ctx.Validate(body); err != nil {
		return err
	}

	return nil
}

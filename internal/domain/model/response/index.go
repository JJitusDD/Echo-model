package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Meta Meta        `json:"meta" mapstructure:"meta"`
	Data interface{} `json:"data,omitempty" mapstructure:"data"`
}

type Meta struct {
	Code     int         `json:"code" mapstructure:"code"`
	Msg      string      `json:"msg" mapstructure:"msg"`
	Message  interface{} `json:"message,omitempty" mapstructure:"message"`
	Internal string      `json:"internal" mapstructure:"internal"`
}

type PaginationResponse struct {
	Page        int64 `json:"page" gorm:"page" mapstructure:"page"`
	TotalItem   int64 `json:"total_item" gorm:"total_item" mapstructure:"total_item"`
	PerPage     int64 `json:"per_page" gorm:"per_page" mapstructure:"per_page"`
	TotalAmount int64 `json:"total_amount,omitempty" mapstructure:"total_amount" gorm:"total_amount"`
}

func NewResponseSuccess(c echo.Context, data interface{}) error {
	if data != nil {
		return c.JSON(http.StatusOK, Response{
			Meta: Meta{
				Code:     http.StatusOK,
				Msg:      http.StatusText(http.StatusOK),
				Internal: http.StatusText(http.StatusOK),
			},
			Data: data,
		})
	}
	return c.JSON(http.StatusNoContent, Response{
		Meta: Meta{
			Code:     http.StatusNoContent,
			Msg:      http.StatusText(http.StatusNoContent),
			Internal: http.StatusText(http.StatusNoContent),
		},
		Data: nil,
	})
}

func NewResponseError(c echo.Context, err error, code int, data interface{}) error {
	if data != nil {
		return c.JSON(http.StatusOK, Response{
			Meta: Meta{
				Code: code,
				Msg:  err.Error(),
			},
			Data: data,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Meta: Meta{
			Code: code,
			Msg:  err.Error(),
		},
		Data: nil,
	})
}

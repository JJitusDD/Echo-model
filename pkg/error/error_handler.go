package error

import (
	"echo-model/pkg/response_definition"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/temporal"
)

func CustomHTTPErrorHandler(l *logrus.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			l.WithError(err).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("echo http error")

			res := response_definition.Response{
				Meta: response_definition.Meta{
					Code:    he.Code,
					Message: he.Message,
				},
				Data: nil,
			}
			if he.Internal != nil {
				res.Meta.Msg = he.Internal.Error()
			}
			return
		} else if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var messages error
			for _, validationErr := range validationErrs {
				// custom handling of specific validation errors
				switch validationErr.Tag() {
				case "required":
					messages = fmt.Errorf("%s is required", validationErr.Field())
				case "email":
					messages = fmt.Errorf("%s must be a valid email", validationErr.Field())
				case "gte":
					messages = fmt.Errorf("%s must be greater than or equal to %s", validationErr.Field(), validationErr.Param())
				case "lte":
					messages = fmt.Errorf("%s must be less than or equal to %s", validationErr.Field(), validationErr.Param())
				default:
					messages = fmt.Errorf("some thing went wrong!")
				}
			}

			l.WithError(err).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("validator error")

			// Return the validation error messages as a JSON response_definition.Response
			c.JSON(http.StatusBadRequest, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadRequest,
					Message: nil,
					Msg:     messages.Error(),
				},
				Data: nil,
			})
			return
		} else if we, ok := err.(*temporal.WorkflowExecutionError); ok {
			l.WithError(we).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("workflow error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     we.Error(),
					Message: we,
				},
				Data: nil,
			})
			return
		} else if ae, ok := err.(*temporal.ActivityError); ok {
			l.WithError(ae).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("activity error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     ae.Error(),
					Message: ae,
				},
				Data: nil,
			})
			return
		} else if se, ok := err.(*temporal.ServerError); ok {
			l.WithError(se).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("server error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     se.Error(),
					Message: se,
				},
				Data: nil,
			})
			return
		} else if te, ok := err.(*temporal.TerminatedError); ok {
			l.WithError(te).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("terminated error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     te.Error(),
					Message: te,
				},
				Data: nil,
			})
			return
		} else if toe, ok := err.(*temporal.TimeoutError); ok {
			l.WithError(toe).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("timeout error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     toe.Error(),
					Message: toe,
				},
				Data: nil,
			})
			return
		} else if ue, ok := err.(*temporal.UnknownExternalWorkflowExecutionError); ok {
			l.WithError(ue).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("unknow workflow error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     ue.Error(),
					Message: ue,
				},
				Data: nil,
			})
			return
		} else if pe, ok := err.(*temporal.PanicError); ok {
			l.WithError(pe).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("unknow workflow error handler")

			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     pe.Error(),
					Message: pe,
				},
				Data: nil,
			})
			return
		} else {
			c.JSON(http.StatusBadGateway, response_definition.Response{
				Meta: response_definition.Meta{
					Code:    http.StatusBadGateway,
					Msg:     err.Error(),
					Message: pe,
				},
				Data: nil,
			})
			return
		}
		// If it's not a validation error, return the default error response
		c.Echo().DefaultHTTPErrorHandler(err, c)

	}
}

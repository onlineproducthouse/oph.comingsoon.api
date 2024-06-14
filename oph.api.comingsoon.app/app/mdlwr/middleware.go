package mdlwr

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/config"
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/loglocal"
	"oph.api.comingsoon/oph.api.comingsoon.util/errorlocal"
	"oph.api.comingsoon/oph.api.comingsoon.util/httpresponse"
)

type MiddlewareAPI struct {
	config config.IConfig
	logger loglocal.ILogger
}

func New(
	config config.IConfig,
	logger loglocal.ILogger,
) MiddlewareAPI {
	return MiddlewareAPI{config, logger}
}

func (api MiddlewareAPI) RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(api.config.RequestIDKey(), uuid.NewString())
		return next(c)
	}
}

func (api MiddlewareAPI) HTTPReqLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		api.logger.HTTPRequest(c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(api.config.RequestIDKey())))
		return next(c)
	}

}

func (api MiddlewareAPI) NotImplemented(env string) func(next echo.HandlerFunc) echo.HandlerFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			if strings.EqualFold(api.config.EnvName(), env) {
				err := errorlocal.NotImplementedErr("MiddlewareAPI.NotImplemented")
				api.logger.AppError(err)
				return c.JSON(err.StatusCode(), httpresponse.Default(err.Error(), err.StatusCode()))
			}

			return next(c)
		}
	}
}

package healthcheckctrl

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/config"
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/loglocal"
	"oph.api.comingsoon/oph.api.comingsoon.service/healthchecksvc"
	"oph.api.comingsoon/oph.api.comingsoon.util/httpresponse"
	"oph.api.comingsoon/oph.api.comingsoon.util/httpstatuslocal"
)

type HealthCheck struct {
	logger loglocal.ILogger
	hc     healthchecksvc.IHealthCheckService
	config config.IConfig
}

type IHealthCheck interface {
	Ping(c echo.Context) error
}

func New(logger loglocal.ILogger, hc healthchecksvc.IHealthCheckService, config config.IConfig) HealthCheck {
	return HealthCheck{logger, hc, config}
}

// HealthCheck/Ping godoc
// @id HealthCheck.Ping
// @tags HealthCheck
// @summary Gets the health status of server.
// @router /HealthCheck/Ping [get]
// @accept */*
// @produce json
// @success 200 {object} httpresponse.Response
// @Failure 500 {object} httpresponse.Response
func (ctrl HealthCheck) Ping(c echo.Context) error {

	if err := ctrl.hc.Ping(); err != nil {
		ctrl.logger.Fatal(err.Error())
	}

	statusCode, statusCodeText := httpstatuslocal.Ok()
	ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))

	return c.JSON(statusCode, httpresponse.Default(statusCodeText, statusCode))
}

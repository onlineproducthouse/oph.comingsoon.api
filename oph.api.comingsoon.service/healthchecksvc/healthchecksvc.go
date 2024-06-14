package healthchecksvc

import (
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/loglocal"
)

type HealthCheckService struct {
	logger loglocal.ILogger
}

type IHealthCheckService interface {
	Ping() error
}

func NewHealthCheckService(logger loglocal.ILogger) HealthCheckService {
	return HealthCheckService{logger}
}

func (svc HealthCheckService) Ping() error {
	const op string = "HealthCheckService.Ping"

	svc.logger.Debug("OK")

	return nil
}

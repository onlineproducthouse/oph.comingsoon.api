package ioc

import (
	infrastructure "oph.api.comingsoon/oph.api.comingsoon.infrastructure"
	"oph.api.comingsoon/oph.api.comingsoon.service/healthchecksvc"
)

type Ioc struct {
	infra infrastructure.IInfrastructure
}

func NewIoc(infra infrastructure.IInfrastructure) Ioc {
	return Ioc{infra}
}

func (ioc Ioc) HealthCheck() healthchecksvc.HealthCheckService {
	return healthchecksvc.NewHealthCheckService(ioc.infra.Logger())
}

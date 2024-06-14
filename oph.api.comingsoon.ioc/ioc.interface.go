package ioc

import "oph.api.comingsoon/oph.api.comingsoon.service/healthchecksvc"

type IIoc interface {
	HealthCheck() healthchecksvc.HealthCheckService
}
